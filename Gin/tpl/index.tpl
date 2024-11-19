<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 0;
            padding: 0;
            background-color: #f4f4f4;
        }
        header {
            background: #333;
            color: #fff;
            padding: 1rem 0;
            text-align: center;
        }
        main {
            padding: 2rem;
        }
        footer {
            background: #333;
            color: #fff;
            text-align: center;
            padding: 1rem 0;
            position: absolute;
            bottom: 0;
            width: 100%;
        }
    </style>
</head>
<body>
    <header>
        <h1>{{.Header}}</h1>
    </header>
    <main>
        <p>{{.Content}}</p>
    </main>
    <footer>
        <p>{{.Footer}}</p>
    </footer>
</body>
</html>
