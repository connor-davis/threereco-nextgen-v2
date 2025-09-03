package services

import (
	"github.com/connor-davis/threereco-nextgen/internal/storage"
)

type Services interface {
	Users() usersService
	Roles() rolesService
	Organizations() organizationsService
	Addresses() addressesService
	BankDetails() bankDetailsService
	Materials() materialsService
	Collections() collectionsService
	Transactions() transactionsService
}

type services struct {
	storage       storage.Storage
	users         usersService
	roles         rolesService
	organizations organizationsService
	addresses     addressesService
	bankDetails   bankDetailsService
	materials     materialsService
	collections   collectionsService
	transactions  transactionsService
}

func NewServices(storage storage.Storage) Services {
	users := newUsersService(storage)
	roles := newRolesService(storage)
	organizations := newOrganizationsService(storage)
	addresses := newAddressesService(storage)
	bankDetails := newBankDetailsService(storage)
	materials := newMaterialsService(storage)
	collections := newCollectionsService(storage)
	transactions := newTransactionsService(storage)

	return &services{
		storage:       storage,
		users:         users,
		roles:         roles,
		organizations: organizations,
		addresses:     addresses,
		bankDetails:   bankDetails,
		materials:     materials,
		collections:   collections,
		transactions:  transactions,
	}
}

func (s *services) Users() usersService {
	return s.users
}

func (s *services) Roles() rolesService {
	return s.roles
}

func (s *services) Organizations() organizationsService {
	return s.organizations
}

func (s *services) Addresses() addressesService {
	return s.addresses
}

func (s *services) BankDetails() bankDetailsService {
	return s.bankDetails
}

func (s *services) Materials() materialsService {
	return s.materials
}

func (s *services) Collections() collectionsService {
	return s.collections
}

func (s *services) Transactions() transactionsService {
	return s.transactions
}
