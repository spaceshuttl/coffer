export function accounts(state = [], action) {
  switch (action.type) {
    case ADD_ACCOUNT:
      return [
        ...state,
        {
          text: action.text,
          completed: false
        }
      ]
    // case DELETE_ACCOUNT:
      // return state.filter((_, i) => i !== action.id),
    default:
      return state
  }
}
