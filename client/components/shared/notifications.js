import m from 'mithril'
import { guid } from '../../utils/helpers'

let state = {
    list: [],
    destroy(msg) {
        let index = state.list.findIndex(x => x.id === msg.id)
        state.list.splice(index, 1)
    }
}

export function addSuccess(text, timeout = 3000) {
    addNotification(guid(), 'success', text, timeout)
}
export function addInfo(text, timeout = 3000) {
    addNotification(guid(), 'info', text, timeout)
}
export function addWarning(text, timeout = 3000) {
    addNotification(guid(), 'warning', text, timeout)
}
export function addDanger(text, timeout = 3000) {
    addNotification(guid(), 'danger', text, timeout)
}

function addNotification(id, type, text, timeout) {
    let msg = { id, type, text, timeout }
    state.list.push(msg)
    setTimeout(() => {
        state.destroy(msg)
        m.redraw()
    }, timeout)
}

const Notifications = {
    oninit(vnode) {
    },
    notificationClass(type) {
        const types = ['info', 'warning', 'success', 'danger']
        if (types.indexOf(type) > -1)
            return type
        return 'info'
    },
    view(vnode) {
        let ui = vnode.state
        return state.list ?
            m('.m-notifications', state.list.map((msg) => {
                return m('.m-notification', { key: msg.id, class: ui.notificationClass(msg.type), onclick: () => { state.destroy(msg) } }, msg.text)
            })) : null
    }
}

export default Notifications