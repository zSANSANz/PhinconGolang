const https = require('https');

https.get('https://pokeapi.co/api/v2/pokedex/kanto/', (resp) => {
    let data = '';

    // A chunk of data has been received.
    resp.on('data', (chunk) => {
    data += chunk;
    });

    // The whole response has been received. Print out the result.
    resp.on('end', () => {
    console.log(JSON.parse(data));
    });

    const express = require('express')
    const app = express()
    const port = 3000

    app.get('/', (req, res) => {
    res.send(JSON.parse(data))
    })

    app.listen(port, () => {
    console.log(`Example app listening at http://localhost:${port}`)
    })

}).on("error", (err) => {
    console.log("Error: " + err.message);
  });