# MongoDB CRUD Application

A full-stack web application built with Go, MongoDB, and Bootstrap that provides a complete CRUD (Create, Read, Update, Delete) interface for user management with advanced features like search functionality, password encryption, and real-time validation.

## 🚀 Features

- **Complete CRUD Operations**: Create, read, update, and delete users with full validation
- **Timestamp Tracking**: Automatic creation and update timestamps for audit trails
- **Email Validation**: Multi-level email format validation (client and server-side)
- **Responsive Design**: Bootstrap 5-based UI that works seamlessly on desktop and mobile
- **REST API**: JSON endpoints for programmatic access


## 🛠 Technology Stack

| Component | Technology |
|-----------|------------|
| **Backend** | Go 1.19+ |
| **Database** | MongoDB |
| **Web Framework** | Gorilla Mux |
| **Frontend** | HTML5, CSS3, JavaScript ES6+ |
| **UI Framework** | Bootstrap 5 |
| **Authentication** | SCRAM-SHA-256 |
| **Icons** | Bootstrap Icons |

## 📋 Prerequisites

Before running this application, ensure you have:

- [Go](https://golang.org/dl/) (version 1.19 or higher)
- [MongoDB](https://www.mongodb.com/try/download/community) (local installation or MongoDB Atlas)
- [Git](https://git-scm.com/downloads)
- Modern web browser (Chrome, Firefox, Safari, Edge)

## 🔧 Installation & Setup

### 1. Clone the Repository
```bash
git clone https://github.com/delacruzjardon/mongodb-crud-app.git
cd mongodb-crud-app
```

### 2. Install Dependencies
```bash
go mod download
```

### 3. Database Configuration

#### Option A: Local MongoDB
```bash
# Start MongoDB service
mongod --replSet local --fork --logappend --logpath ~/db/logs/mongod.log --dbpath ~/db/data/
```

#### Option B: MongoDB Atlas
1. Create a free cluster at [MongoDB Atlas](https://www.mongodb.com/cloud/atlas)
2. Get your connection string
3. Update `main.go` with your Atlas connection string

### 4. Application Configuration

Update the MongoDB connection in `main.go`:

```go
// Local MongoDB
connectionString := "mongodb://localhost:27017/crudapp"

// MongoDB Atlas with authentication
connectionString := "mongodb://username:password@cluster.mongodb.net/crudapp?authSource=admin&authMechanism=SCRAM-SHA-256"
```

### 5. Environment Variables (Optional)
```bash
export MONGODB_URI="mongodb://localhost:27017/crudapp"
export PORT="8080"
```

## 🚀 Running the Application

```bash
# Start the server
go run main.go

# Or build and run
go build -o mongodb-crud-app
./mongodb-crud-app
```

**Access the application**: Open your browser and navigate to `http://localhost:8080`

## 📁 Project Structure

```
mongodb-crud-app/
├── 📄 main.go                   # Application entry point & server configuration
├── 📁 models/
│   └── 📄 user.go              # User data model & structures
├── 📁 handlers/
│   └── 📄 user_handlers.go     # HTTP handlers for all CRUD operations
├── 📁 utils/
│   ├── 📄 password.go          # Password encryption & validation utilities
│   └── 📄 validation.go        # Email & input validation utilities
├── 📁 templates/
│   ├── 📄 layout.html          # Base HTML template with navigation
│   ├── 📄 index.html           # Home page with user listing
│   ├── 📄 create.html          # Create user form
│   ├── 📄 edit.html            # Edit user form
│   └── 📄 search.html          # Search results page
├── 📁 static/
│   ├── 📁 css/
│   │   └── 📄 style.css        # Custom styles & responsive design
│   └── 📁 js/
│       └── 📄 script.js        # Client-side JavaScript & validation
├── 📄 go.mod                   # Go modules dependencies
├── 📄 go.sum                   # Dependency checksums
└── 📄 README.md               # This documentation
```

## 🔗 API Endpoints

### Web Routes
| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/` | Home page with user list |
| `GET` | `/create` | Create user form |
| `POST` | `/create` | Create new user |
| `GET` | `/edit/{id}` | Edit user form |
| `POST` | `/update/{id}` | Update existing user |
| `GET` | `/delete/{id}` | Delete user |
| `GET` | `/search` | Search users by name/email |

### API Routes (JSON)
| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/api/users` | Get all users (JSON response) |
| `GET` | `/api/search?q={query}` | Search users (JSON response) |

### Example API Usage
```bash
# Get all users
curl http://localhost:8080/api/users

# Search users
curl "http://localhost:8080/api/search?q=john"

# Expected JSON response format
{
  "_id": "507f1f77bcf86cd799439011",
  "name": "John Doe",
  "email": "john.doe@example.com",
  "age": 30,
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

## 🔒 Security Features

### Password Security
- **bcrypt Hashing**: Industry-standard password hashing with salt
- **SCRAM-SHA Compatibility**: Compatible with MongoDB's authentication
- **Minimum Requirements**: 8+ characters, complexity validation
- **Secure Comparison**: Constant-time comparison to prevent timing attacks

### Input Validation
- **Email Validation**: RFC-compliant email format checking
- **Server-side Validation**: All inputs validated on the server
- **Client-side Validation**: Real-time feedback for better UX
- **XSS Prevention**: Input sanitization and escaping
- **SQL Injection Prevention**: MongoDB driver handles query sanitization

### Data Protection
- **Password Fields**: Never exposed in JSON responses
- **Secure Headers**: CSRF protection and secure headers
- **Connection Security**: TLS/SSL support for MongoDB connections

## 🎨 User Interface Features

### Design & UX
- **Responsive Layout**: Mobile-first design with Bootstrap 5
- **Accessibility**: ARIA labels, semantic HTML, keyboard navigation
- **Visual Feedback**: Success/error states, loading indicators
- **Modern UI**: Clean, professional interface with consistent styling

### Interactive Features
- **Real-time Search**: Live search with autocomplete suggestions
- **Form Validation**: Immediate feedback on form inputs
- **Keyboard Shortcuts**: `Ctrl+K` or `Cmd+K` to focus search
- **Confirmation Dialogs**: Safe deletion with user confirmation
- **Error Handling**: User-friendly error messages

## 📊 Usage Examples

### Creating a New User
1. Navigate to the home page (`http://localhost:8080`)
2. Click "Add New User" button
3. Fill in the required information:
   - **Name**: Required, 1-100 characters
   - **Email**: Required, valid email format
   - **Password**: Required, minimum 8 characters
   - **Age**: Required, 1-120 years
4. Click "Create User" to save

### Searching for Users
1. Use the search box on any page
2. Type name or email to search
3. Get real-time suggestions as you type
4. Press Enter or click "Search" for full results
5. View filtered results with match highlighting

### Updating User Information
1. Find the user in the list or search results
2. Click the "Edit" button
3. Modify the information (password field is optional)
4. Click "Update User" to save changes

## 🧪 Testing

### Running Tests
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./handlers
go test ./utils
```

### Manual Testing Checklist
- [ ] Create user with valid data
- [ ] Create user with invalid email
- [ ] Create user with weak password
- [ ] Update existing user
- [ ] Delete user (with confirmation)
- [ ] Search functionality
- [ ] Responsive design on mobile
- [ ] Form validation (client & server)

## ⚡ Performance Considerations

### Database Optimization
- **Email Indexing**: Email field indexed for faster searches
- **Connection Pooling**: MongoDB driver handles connection pooling
- **Query Optimization**: Efficient regex queries for search
- **Pagination**: Ready for implementation as data grows

### Frontend Optimization
- **Static File Caching**: CSS/JS files served with cache headers
- **Minification**: Production-ready CSS and JS
- **Lazy Loading**: Images and non-critical resources
- **CDN Usage**: Bootstrap and icons served from CDN

## 🔧 Configuration Options

### Environment Variables
```bash
# Database configuration
export MONGODB_URI="mongodb://localhost:27017/crudapp"
export DATABASE_NAME="crudapp"
export COLLECTION_NAME="users"

# Server configuration
export PORT="8080"
export HOST="localhost"

# Security configuration
export BCRYPT_COST="12"
export SESSION_SECRET="your-secret-key"
```

### MongoDB Schema Validation
The application includes optional MongoDB schema validation:

```javascript
// Example schema validation
{
  "bsonType": "object",
  "required": ["name", "email", "age"],
  "properties": {
    "name": {
      "bsonType": "string",
      "minLength": 1,
      "maxLength": 100
    },
    "email": {
      "bsonType": "string",
      "pattern": "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"
    },
    "age": {
      "bsonType": "int",
      "minimum": 1,
      "maximum": 120
    }
  }
}
```

## 🤝 Contributing

We welcome contributions! Please follow these steps:

1. **Fork the repository**
2. **Create a feature branch**:
   ```bash
   git checkout -b feature/amazing-feature
   ```
3. **Make your changes**:
   - Add tests for new functionality
   - Update documentation as needed
   - Follow Go coding standards
4. **Commit your changes**:
   ```bash
   git commit -m 'Add some amazing feature'
   ```
5. **Push to your branch**:
   ```bash
   git push origin feature/amazing-feature
   ```
6. **Open a Pull Request**

### Contribution Guidelines
- Write clear, concise commit messages
- Include tests for new features
- Update README if adding new functionality
- Follow existing code style and conventions
- Add comments for complex logic

## 🆘 Troubleshooting

### Common Issues & Solutions

#### MongoDB Connection Error
```
Error: connection refused
```
**Solutions:**
- Ensure MongoDB is running: `mongod --version`
- Check connection string format
- Verify network connectivity and firewall settings
- For Atlas: Check IP whitelist and credentials

#### Port Already in Use
```
Error: bind: address already in use
```
**Solutions:**
```bash
# Find process using port 8080
lsof -i :8080

# Kill the process
kill -9 <PID>

# Or change port in main.go
```

#### Import Path Errors
```
Error: package not found
```
**Solutions:**
```bash
# Ensure you're in project directory
pwd

# Download dependencies
go mod download

# Clean module cache if needed
go clean -modcache
```

#### Template Not Found
```
Error: template not found
```
**Solutions:**
- Ensure `templates/` directory exists
- Check file permissions
- Verify template file names match handler references

## 📚 Additional Resources

### MongoDB Resources
- [MongoDB Documentation](https://docs.mongodb.com/)
- [MongoDB University](https://university.mongodb.com/) - Free courses
- [MongoDB Atlas](https://www.mongodb.com/cloud/atlas) - Cloud database service
- [MongoDB Community Forums](https://community.mongodb.com/)

### Go Resources
- [Go Documentation](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go by Example](https://gobyexample.com/)
- [Gorilla Mux Documentation](https://github.com/gorilla/mux)

### Frontend Resources
- [Bootstrap Documentation](https://getbootstrap.com/docs/)
- [MDN Web Docs](https://developer.mozilla.org/)
- [CSS-Tricks](https://css-tricks.com/)

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

```
MIT License

Copyright (c) 2024 delacruzjardon

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
```

## 🙏 Acknowledgments

- **MongoDB Team** for the excellent database and documentation
- **Bootstrap Team** for the responsive CSS framework
- **Go Community** for the powerful programming language and ecosystem
- **Open Source Contributors** for the various libraries used in this project

## 📞 Support & Contact

- **Issues**: [GitHub Issues](https://github.com/delacruzjardon/mongodb-crud-app/issues)
- **Discussions**: [GitHub Discussions](https://github.com/delacruzjardon/mongodb-crud-app/discussions)
- **Maintainer**: [@delacruzjardon](https://github.com/delacruzjardon)

---

**⭐ If this project helped you, please consider giving it a star on GitHub!**

**Built with ❤️ using Go, MongoDB, and modern web technologies**

