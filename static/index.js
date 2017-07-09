const wsuri = "ws://127.0.0.1:1234/ws"

const checkIfUrlOrFilepath = (imageSrc) => {
  if (imageSrc.indexOf('http') > -1) {
    return e.data
  } else {
    return "./images/" + imageSrc
  }
}

window.onload = () => {
  const sock = new WebSocket(wsuri);

  sock.onmessage = (e) => {
    const link = checkIfUrlOrFilepath(e.data)
    document.getElementById("img").style.backgroundImage = 'url("' + link + '")'
  }
}