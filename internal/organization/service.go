// Package organization owns organization creation independently of authentication providers.
package organization

import (
	"context"
	"net/mail"
	"strings"
	"unicode/utf8"
)

type Service struct{ repository Repository }

func NewService(
	repository Repository,
) *Service {
	return &Service{repository: repository}
}

func (s *Service) Create(
	ctx context.Context,
	identity Identity,
	name, email string,
) (Organization, error) {
	name, email = strings.TrimSpace(name), strings.TrimSpace(email)
	var fields []FieldError
	if name == "" {
		fields = append(fields, FieldError{"name", "required"})
	} else if strings.ContainsRune(name, '\x00') {
		fields = append(fields, FieldError{"name", "invalid"})
	} else if utf8.RuneCountInString(name) > maxNameRunes {
		fields = append(fields, FieldError{"name", "too_long"})
	}
	if email == "" {
		fields = append(
			fields,
			FieldError{"initialAdministratorEmail", "required"},
		)
	} else if len(email) > maxEmailBytes {
		fields = append(
			fields,
			FieldError{"initialAdministratorEmail", "too_long"},
		)
	} else {
		address, err := mail.ParseAddress(email)
		// ParseAddress also accepts display names and comments; require the bare address verbatim.
		if err != nil || address.Name != "" || address.Address != email ||
			strings.ContainsAny(email, "\r\n") {
			fields = append(
				fields,
				FieldError{"initialAdministratorEmail", "invalid"},
			)
		} else if at := strings.LastIndexByte(
			email,
			'@',
		); at < 1 ||
			at == len(email)-1 {
			fields = append(
				fields,
				FieldError{"initialAdministratorEmail", "invalid"},
			)
		} else {
			email = email[:at+1] + strings.ToLower(email[at+1:])
		}
	}
	if len(fields) > 0 {
		return Organization{}, &ValidationError{fields}
	}
	if identity.Issuer == "" || identity.Subject == "" {
		return Organization{}, ErrForbidden
	}
	return s.repository.Create(ctx, identity, name, email)
}
