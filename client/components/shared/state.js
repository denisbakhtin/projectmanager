import {
    responseErrors
} from '../../utils/helpers'
import service from '../../utils/service.js'

//global ui state
const state = {
    sidebarCollapsed: false,
    notifications: [],
    errors: [],
    favoriteProjects: [],
    query: '',

    //methods
    toggleSidebar: () => state.sidebarCollapsed = !state.sidebarCollapsed,
    setQuery: (val) => state.query = val,

    getNotifications: () =>
        service.getNotifications()
            .then((result) => state.notifications = result.slice(0))
            .catch((error) => state.errors = responseErrors(error)),

    removeNotification: (id) => service.deleteNotification(id),

    getFavoriteProjects: () =>
        service.getFavoriteProjects()
            .then((result) => state.favoriteProjects = result.slice(0))
            .catch((error) => state.errors = responseErrors(error))
}

export default state
