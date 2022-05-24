const fastify = require('fastify')({
  logger: true
})

fastify.get('/', async (request, reply) => {
  return { hello: 'Docker Meetup' }
})

const start = async () => {
  try {
    await fastify.listen(3000, '0.0.0.0')
  } catch (err) {
    fastify.log.error(err)
    process.exit(1)
  }

  process.on('SIGTERM', async () => {
    console.log('Terminating...');
    await fastify.close()
    process.exit(0)
  })

  process.on('SIGINT', async () => {
    console.log('Interrupt...');
    await fastify.close()
    process.exit(0)
  })
}

start()