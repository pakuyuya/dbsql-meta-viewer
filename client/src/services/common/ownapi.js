'use struct'

import config from 'config'

export default {
  resolveurl: function (path) {
    return config.apiurl
      .replace(/\{protocol\}/, document.location.protocol)
      .replace(/\{host\}/, document.location.host) + path
  }
}
