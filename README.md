# GoNexus - E-commerce Platform

GoNexus is a modern e-commerce platform built with Go, featuring a beautiful UI and robust functionality. The platform supports user authentication, product management, and a responsive shopping experience.

## Features

- 🛍️ Modern E-commerce UI with Tailwind CSS
- 🔐 Secure Authentication System
  - JWT-based authentication
  - OAuth2 support (Google, GitHub)
  - Password reset functionality
  - Email verification
- 🎨 Responsive Design
- 🛒 Shopping Cart System
- 📱 Mobile-friendly Interface
- 🔍 Product Search and Filtering
- 📦 Product Categories
- 💳 Payment Integration (coming soon)

## Tech Stack

- **Backend**: Go
- **Frontend**: 
  - Templ (Go template engine)
  - Tailwind CSS
  - HTMX for dynamic interactions
- **Database**: PostgreSQL
- **Authentication**: JWT, OAuth2
- **Email Service**: SMTP
- **File Storage**: Local (configurable)

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 15 or higher
- Node.js and npm (for Tailwind CSS)

## Environment Variables

Create a `.env` file in the root directory with the following variables:

```env
# Server Configuration
PORT=8080
ENV=development

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=gonexus

# JWT Configuration
JWT_SECRET=your_jwt_secret
JWT_EXPIRATION=24h

# OAuth2 Configuration
GOOGLE_CLIENT_ID=your_google_client_id
GOOGLE_CLIENT_SECRET=your_google_client_secret
GITHUB_CLIENT_ID=your_github_client_id
GITHUB_CLIENT_SECRET=your_github_client_secret

# Email Configuration
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your_email@gmail.com
SMTP_PASSWORD=your_app_specific_password
SMTP_FROM_EMAIL=your_email@gmail.com
SMTP_FROM_NAME=GoNexus

# File Storage
UPLOAD_DIR=./uploads
MAX_UPLOAD_SIZE=5242880  # 5MB in bytes
```

## Installation

1. Clone the repository:
```bash
git clone https://github.com/jimtrung/go-nexus.git
cd go-nexus
```

2. Install dependencies:
```bash
go mod download
```

3. Set up the database:
```bash
# Create the database
createdb gonexus

# Run migrations (if available)
go run cmd/migrate/main.go
```

4. Configure environment variables:
```bash
cp .env.example .env
# Edit .env with your configuration
```

5. Build and run the application:
```bash
go run cmd/main.go
```

## Development

### Frontend Development

The project uses Tailwind CSS for styling. To compile CSS:

```bash
# Install Tailwind CSS
npm install

# Watch for changes
npm run watch
```

### Project Structure

```
gonexus/
├── cmd/                    # Application entry points
├── internal/              # Private application code
│   ├── api/              # API handlers and routes
│   ├── domain/           # Business logic and models
│   ├── infra/            # Infrastructure code
│   ├── repository/       # Database repositories
│   └── services/         # Business services
├── templates/            # HTML templates
│   ├── component/        # Reusable components
│   └── layout/          # Layout templates
├── static/              # Static assets
├── uploads/             # User uploaded files
└── go.mod              # Go module file
```

## API Endpoints

### Authentication
- `POST /auth/signup` - User registration
- `POST /auth/login` - User login
- `POST /auth/logout` - User logout
- `POST /auth/forgot-password` - Request password reset
- `POST /auth/reset-password` - Reset password
- `GET /auth/verify/:token` - Verify email

### Products
- `GET /products` - List products
- `GET /products/:id` - Get product details
- `POST /products` - Create product (admin)
- `PUT /products/:id` - Update product (admin)
- `DELETE /products/:id` - Delete product (admin)

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- [Tailwind CSS](https://tailwindcss.com/) for the styling
- [Templ](https://templ.guide/) for the template engine
- [HTMX](https://htmx.org/) for dynamic interactions
