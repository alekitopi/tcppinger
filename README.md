# TCPPinger

> Comprueba la disponibilidad y tiempo de tramisión de paquetes dentro de la red en un puerto TCP determinado.

Con esta herramienta podrás comprobar la disponibilidad de varios servicios como también conocer el tiempo exacto que tarda en transferirse un paquete a dicho servicio.

---

## Uso

```bash
./pinger <ip> <puerto> [<opciones>]
```

### Opciones

| Opción   | Descripción                          |
| -------- | ------------------------------------ |
| -h       | Muestra el mensaje de ayuda.         |
| -t <seg> | Modifica el delay entre cada pingeo. |

## Ejemplo

```bash
./pinger localhost 80 -t 1
```

## Compilar

Para realizar alguna modificación del código y compilar el mismo deberás tener instalado Go y Git en tu sistema operativo.
(en este caso utilizaremos Linux)

```bash
sudo apt-get install git golang -y
git clone https://github.com/alekitopi/tcppinger.git
cd tcppinger/
go build ./pinger.go
```
