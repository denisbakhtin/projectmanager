import m from 'mithril';
import {
  responseErrors
} from '../../utils/helpers'
import Auth from '../../utils/auth'

//global ui state
const state = {
  sidebarCollapsed: false,
  notifications: [],
  errors: [],

  //methods
  toggleSidebar() {
    state.sidebarCollapsed = !state.sidebarCollapsed
  },
  getNotifications() {
    m.request({
        method: "GET",
        url: "/api/notifications",
        headers: {
          Authorization: Auth.authHeader()
        },
      })
      .then((result) => {
        state.notifications = result.slice(0)
      })
      .catch((error) => state.errors = responseErrors(error))
  },
  removeNotification(id) {
    return m.request({
      method: "DELETE",
      url: "/api/notifications/" + id,
      headers: {
        Authorization: Auth.authHeader()
      }
    });
  }
}

export default state