# Docpst

[![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Status: WIP](https://img.shields.io/badge/Status-Active%20Development-orange.svg)](https://github.com/alexxbez/docpst)

A powerful Go CLI tool for managing Typst projects. Create, organize, and manage your Typst documents and templates with ease.

## 🚧 Project Status

> **Note**: This project is in active development and not yet ready for production use. APIs may change without notice.

## 📖 Description

Docpst is a command-line interface tool written in Go that simplifies Typst project management. It helps you:

- **Create** new Typst projects
- **Organize** your documents and templates efficiently  
- **Edit** documents with your preferred workflow
- **Manage** templates across multiple projects

## 🚀 Quick Start

### Prerequisites

- **[Go](https://golang.org/dl/)** 1.21 or later (for development)
- **[Typst](https://github.com/typst/typst)** installed and in your PATH
- **[Zellij](https://zellij.dev/)** (for terminal workspace management)
- **[Zathura](https://pwmt.org/projects/zathura/)** (for PDF preview)

### Installation

```bash
go install github.com/alexxbez/docpst/cmd/docpst@latest
```

### Your First Command

```bash
docpst --help
```

## 💡 Usage Examples

```bash
# Working commands
docpst new my-document          # Create a new Typst project
docpst watch                    # Compile Typst document
```

## 🛠️ Development

### Building from Source

```bash
git clone https://github.com/alexxbez/docpst.git
cd docpst
go build -o docpst ./cmd/docpst
```

## 📝 Version History

### v0.1.0
- Project skeleton and basic structure
- Initial CLI framework
- Basic command setup

## 🤝 Contributing

The project is not ready to recieve contributions yet. Once docpst reaches a stable version contributions will begin.

## 👤 Author

**Alejandro Benítez Bravo**
- Email: [alexxbez@proton.me](mailto:alexxbez@proton.me)
- GitHub: [@alexxbez](https://github.com/alexxbez)

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.
