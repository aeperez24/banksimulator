# banksimulator

Mini Banco
Se requiere crear una aplicación web que simule las funcionalidades de un banco, para ello se
necesitan los siguientes módulos:
Home con acción de registro y login.
Registro
Formulario que permita crear una nueva cuenta, los campos serian:
Nombre, Rut, correo y contraseña. Todos los campos son obligatorios, solo debe permitir una cuenta
por Rut.
Carga de Saldo
Formulario que permita agregar fondos una cuenta (simulando un depósito de fondos), solo debe
llevar un input que permita ingresar el monto a depositar y un botón que permita aceptar.
Retiro de Saldo
Formulario que permita retirar dinero de una cuenta, solo un campo del monto a retirar y un aceptar.
La cuenta no puede quedar con saldo negativo
Transferencia
Formulario para transferir a un tercero, se debe solicitar como entrada el Rut destino y el monto a
transferir, el monto permitido debe ser menor o igual al saldo disponible en la cuenta origen. La
cuenta origen no puede quedar con saldo negativo, se debe validar que la cuenta destino este
registrada en el sistema.
Listado de movimientos
Debe mostrar los movimientos realizados (para cada cliente consultado), ya se cargas de saldo y
transferencias (entrantes y salientes).
