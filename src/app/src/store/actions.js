const ADD_ACCOUNT = "ADD"
const DELETE_ACCOUNT = "DELETE"

function addAccount(account) {
  return {
    type: ADD_ACCOUNT,
    account,
  }
}

function deleteAccount(account) {
  return {
    type: DELETE_ACCOUNT,
    account,
  }
}
