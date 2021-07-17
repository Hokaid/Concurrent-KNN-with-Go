<?php
header('Access-Control-Allow-Origin: http://localhost:9000/');
?>
<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.1/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-+0n0xVW2eSR5OomGNYDnhzAbDsOXxcvSN1TPprVMTNDbiYZCxYbOOl7+AMvyTG2x" crossorigin="anonymous">
    <link href="https://unpkg.com/tailwindcss@^2/dist/tailwind.min.css" rel="stylesheet">
    <title>TA2</title>
</head>
<body class="bg-blue-100">
    <div class="text-center mt-6">
        <h1 class="text-xl font-bold">Aplicación de KNN en un dataset de Transporte</h1>
    </div>
    <div class="flex justify-center mt-16">
        <div class="lg:w-5/12 w-full mx-3">
            <div class="form-floating mb-4">
                <input type="number" class="form-control" id="kinput" aria-describedby="khelp" placeholder="Ingrese un numero entero">
                <label for="kinput">Valor de K</label>
                <div id="khelp" class="form-text">Valor de k a considerar como parametro en el algoritmo KNN.</div>
            </div>
            <div class="form-floating mb-4">
                <input type="number" step="0.01" class="form-control" id="ptestinput" aria-describedby="ptesthelp" placeholder="Ingrese un numero del 0 al 1">
                <label for="ptestinput">Porcentaje de Test</label>
                <div id="ptesthelp" class="form-text">Indicar el porcentaje de datos del dataset que se van a considerar para test.</div>
            </div>
        </div>
    </div>
    <div class="flex justify-center mt-11">
        <button onclick="ObtenerResultados()" type="button" class="btn btn-info">Obtener Resultados</button>
    </div>
    <div class="flex justify-center mt-10">
        <div id="resultado" class="text-lg font-bold lg:w-2/6 w-full mx-3 text-center">Ingrese los parametros y seleccione el boton para obtener resultados</div>
    </div>
    <script type="text/javascript">
        const ObtenerResultados = async () => {
            var k = document.getElementById("kinput").value;
            var ptest = document.getElementById("ptestinput").value;
            const response = await fetch('http://localhost:9000/test?k=' + k.toString() + '&ptest=' + ptest.toString());
            const myJson = await response.json();
            document.getElementById("resultado").innerHTML = "Exactitud Obtenida: "+(myJson.acc*100) + "%";
        }
    </script>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.1/dist/js/bootstrap.bundle.min.js" integrity="sha384-gtEjrD/SeCtmISkJkNUaaKMoLD0//ElJ19smozuHV6z3Iehds+3Ulb9Bn9Plx0x4" crossorigin="anonymous"></script>
</body>

</html>