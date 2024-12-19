# Verwendung

## 1. Starte das erste Programm:


```bash
go run prog1.go
```

Es gibt die PID und die Speicheradresse der Variable `value` aus, z. B.:

```yaml
PID: 1234
Adresse von 'value': 0x12345678
```

## 2. Starte das zweite Programm mit PID und Adresse:

```bash
go run prog2.go <PID> <Adresse>

```

Beispiel:

```bash
go run prog2.go 1234 0x12345678

```