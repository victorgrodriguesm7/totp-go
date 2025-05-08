# TOTP Redirector

A small HTTP server written in Go that redirects requests with `issuer` and `secret` parameters to an `otpauth://` URL.

### 🔗 Example usage

Access a URL like this:

```
http://localhost:8080/otp?issuer=MyCompany&secret=ABCDEF123456
```

The server will redirect to:

```
otpauth://totp/MyCompany?secret=ABCDEF123456&issuer=MyCompany
```

> ⚠️ Note: Most browsers do not handle `otpauth://` links directly. This is mainly useful for mobile apps or QR code scanners that support the TOTP scheme.

---

### 🚀 How to run

1. Install Go: https://golang.org/dl/
2. Clone this repository
3. Run the server:

```bash
go run main.go
```

The server will be available at http://localhost:8080.

### 🛠️ Possible improvements

- [ ] Parameter validation

- [ ] QR code generation via HTML page

- [ ] Support for other schemes (e.g. hotp)

- [X] Web UI to simplify link creation

## :page_facing_up: License
This project is licensed under the MIT License - see the [LICENSE](/LICENSE) file for details.

## :mailbox_with_mail: Get in touch!

<a href="mailto:victorgrodriguesm7@gmail.com" target="_blank" >
  <img alt="Email - Victor Gabriel" src="https://img.shields.io/badge/Email--%23F8952D?style=social&logo=gmail">
</a>
<a href="https://www.linkedin.com/in/victorgrodriguesm7/">
    <img src="https://img.shields.io/badge/Linkedin--%23F8952D?style=social&logo=linkedin">
</a>