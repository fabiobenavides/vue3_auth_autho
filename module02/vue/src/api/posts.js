import axios from 'axios';

export default {
  findMany() {
    const config = {};

    return new Promise((resolve, reject) => {
      axios.get("posts", config)
      .then(result => {
        resolve(result.data)
      })
      .catch(error => {
        reject(error)
      })
    })
  }
}
