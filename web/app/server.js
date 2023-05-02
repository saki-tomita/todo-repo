const http = require('http')

if (require.main === module) {
    main()
}

async function main () {
    try {
        const server = http.createServer() // <1>

        server.on('listening', () => { // <2>
            console.info(`Listening on ${process.env.PORT}`)
        })

        server.on('request', (req, res) => { // <3>
            const content = JSON.stringify({ok: true})

            res.writeHead(200, {
                'Content-Type': 'application/json; charset=utf-8',
                'Content-Length': content.length,
            })

            res.write(content)
            res.end()
        })

        server.on('error', err => { // <4>
            console.error(err)
        })

        server.listen(process.env.PORT) // <5>
    } catch (err) {
        console.error(err)
    }
}