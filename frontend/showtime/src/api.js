const fetch = window.fetch

const apiPrefix = 'http://localhost:8081'

const search = (q) => {
  return fetch(`${apiPrefix}/search?q=${encodeURIComponent(q)}`)
    .then((resp) => resp.json())
}

const getResource = (id) => {
  return fetch(`${apiPrefix}/resources/${id}`)
    .then((resp) => resp.json())
}

export default {
  search,
  getResource
}
