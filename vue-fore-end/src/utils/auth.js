const TokenKey = 'user_token'

export function getToken() {
  return localStorage.getItem(TokenKey)
}

export function setToken(token) {
  console.log("setToken", token)
  return localStorage.setItem(TokenKey, token)
}

export function removeToken() {
  return localStorage.removeItem(TokenKey)
}
