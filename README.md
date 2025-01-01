Ein einfacher URL-Kürzungsdienst in Go, der URLs kürzt und Statistiken über deren Nutzung bietet. Dieses Projekt beinhaltet Benutzerregistrierung und -anmeldung, URL-Verfall und Zugriffszählung.

## Features ✨

- Benutzerregistrierung und -anmeldung 🔐
- URLs kürzen und verwalten 🔗
- URL-Verfall nach einer bestimmten Zeit ⏲️
- Zugriffszählung 📊

## Geplante Features 🔮

- Benutzer-Dashboard zur Verwaltung von URLs 🖥️
- Erweiterte Statistiken (z.B. Klicks pro Tag) 📈
- Unterstützung für benutzerdefinierte Kurz-URLs 🎨
- QR-Code-Generierung für gekürzte URLs 📱

## Installation und Ausführung 🚀

1. **Repository klonen**:
   ```sh
   git clone https://github.com/oleeeedev/url-shortener.git
   cd url-shortener
   ```

2. **Abhängigkeiten installieren**:
   ```sh
   go mod tidy
   ```

3. **Server starten**:
   ```sh
   go run main.go
   ```

## Endpunkte 📡

### Benutzerregistrierung

- **URL**: `/register`
- **Methode**: `POST`
- **Body**:
  ```json
  {
    "username": "deinBenutzername",
    "password": "deinPasswort"
  }
  ```

### Benutzeranmeldung

- **URL**: `/login`
- **Methode**: `POST`
- **Body**:
  ```json
  {
    "username": "deinBenutzername",
    "password": "deinPasswort"
  }
  ```

### URL kürzen

- **URL**: `/shorten`
- **Methode**: `POST`
- **Body**:
  ```json
  {
    "url": "https://example.com",
    "expiry_in": 10,  // Ablaufzeit in Minuten
    "custom_url": "optionalCustomURL"  // Optional
  }
  ```

### URL-Weiterleitung

- **URL**: `{shortURL}`
- **Methode**: `GET`
- **Beschreibung**: Leitet zur Original-URL weiter.

### URL-Statistiken abrufen

- **URL**: `/stats/{shortURL}`
- **Methode**: \`GET\`
- **Beschreibung**: Gibt die Statistiken für die gekürzte URL zurück.

## Beispielaufrufe mit Postman 🛠️

### Benutzer registrieren

1. **Neue Anfrage erstellen**:
   - Methode: `POST`
   - URL: `http://localhost:8080/register`
   - Header: `Content-Type: application/json`
   - Body:
     ```json
     {
       "username": "testuser",
       "password": "password123"
     }
     ```

### Benutzer anmelden

1. **Neue Anfrage erstellen**:
   - Methode: `POST`
   - URL: `http://localhost:8080/login`
   - Header: `Content-Type: application/json`
   - Body:
     ```json
     {
       "username": "testuser",
       "password": "password123"
     }
     ```

### URL kürzen

1. **Neue Anfrage erstellen**:
   - Methode: `POST`
   - URL: `http://localhost:8080/shorten`
   - Header: `Content-Type: application/json`
   - Body:
     ```json
     {
       "url\": "https://example.com",
       "expiry_in": 10,
       "custom_url": "example"
     }
     ```

## Lizenz 📜

Dieses Projekt ist unter der MIT-Lizenz lizenziert. Siehe die [LICENSE](LICENSE)-Datei für weitere Details.

## Mitwirkende ✨

- [oleeeedev](https://github.com/oleeeedev)

---
