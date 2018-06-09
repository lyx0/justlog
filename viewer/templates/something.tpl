<html>
    <body>
        <h1>Hello {{index . "name"}}</h1>

        <p>{{ with $x := index . "reverse" }}
           {{ call $x "foobar" }} &lt;-- this will call the $x with parameter "foobar"
           {{ end }}
        </p>

        <h1> FUCK YOU </h1>
    </body>
</html>