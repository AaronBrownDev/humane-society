# Humane Society Management System [ARCHIVED]

> **Collaborative academic project** for Database Implementation (Spring 2025). Archived as coursework is complete.

## Project Contributions

**Aaron (Backend & Database)**
- Database schema design and implementation (15+ tables, migrations, stored procedures)
- Go REST API with repository pattern and JWT authentication
- Docker containerization and database tooling

**[Courtnae James](https://github.com/courtnaejames1) (Frontend)**
- React application with Material-UI components
- User interface for adoption, volunteer, and surrender workflows
- Admin dashboard and form management
- Collaborated on initial ERD design

## Database Implementation Highlights

**Schema Design**
- 15+ tables with proper normalization and foreign key constraints
- Multi-schema organization (`shelter`, `medical`, `people`, `auth`, `audit`)
- Complex relationships (person → adopter/volunteer/owner inheritance pattern)
- Initial ERD designed collaboratively; implementation and iteration by Aaron

**Migration System**
- 56 sequential migrations for version-controlled schema changes
- Idempotent up/down migrations for safe rollbacks
- Stored procedures for complex business logic (adoption workflow, user creation)

**Production Patterns**
- Audit logging via triggers for change tracking
- Row-level security with database roles (Admin, Staff, Volunteer, ReadOnly)
- Proper indexing on foreign keys and frequently-queried columns

**API Layer**
- Go REST API using repository pattern for testable data access
- JWT authentication with refresh token rotation
- Database connection pooling and transaction management

## Tech Stack
**Backend:** Go • SQL Server • Docker • JWT • REST API  
**Frontend:** React • Vite • Material-UI • React Router

---

*This project focused on database design and backend API development*
