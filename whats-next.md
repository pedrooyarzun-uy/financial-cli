# Status actual

## Que sigue?

- Parametrizar obtencion de cuenta.
- 1. Endpoint desde backend - DONE
- 2. Desplegable de cuentas (dropdown_account) - DONE
- Parametrizar currencies.
- 1. Endpoint desde backend - DONE
- 2. Service en frontend de currencies - DONE
- 3. Desplegable de currencies - DONE
- Parametrizar types. - DONE
- 1. Endpoint desde backend - DONE
- 2. Service de frontend de types - DONE
- 3. Desplegable de types - DONE

- Agregar tipo adjustment para corregir desbalances entre app y banco
- 1. Catalogo de tipos - DONE
- 2. Agregar implementacion en el backend - DONE
- 3. Servicio de frontend - DONE
- 4. Implementar en desplegable - DONE

- Implemenentacion de tarjetas de credito
- 1. Crear tabla para tarjetas de credito (
    type CreditCard struct {
        Id         int
        Name       string
        BankID     int
        CurrencyID int
        OwnerID    int
        CloseDay   int
        DueDay     int
        Limit      float64
        CreatedAt  time.Time
    }
)

- 2. Crear repositorios y servicios para poder agregar una tarjeta.
- 3. Ajustar tabla de transacciones para poder agregar gastos de tarjetas de credito
- 4. Pasar a una vista la traida de informacion de las transactions