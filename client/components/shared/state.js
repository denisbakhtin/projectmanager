import m from 'mithril';
import {
    responseErrors
} from '../../utils/helpers'
import Auth from '../../utils/auth'
import service from '../../utils/service.js'

//global ui state
const state = {
    sidebarCollapsed: false,
    notifications: [],
    errors: [],

    //methods
    toggleSidebar: () =>
        state.sidebarCollapsed = !state.sidebarCollapsed,

    getNotifications: () =>
        service.getNotifications()
        .then((result) => {
            state.notifications = result.slice(0)
        })
        .catch((error) => state.errors = responseErrors(error)),

    removeNotification: (id) => service.deleteNotification(id)
}

export default state
