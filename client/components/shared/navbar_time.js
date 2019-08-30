import m from 'mithril'
import moment from 'moment'

const format = 'MMMM Do YYYY, h:mm a'
let state = {
    now: '',
    timer: null,
    update: () => {
        state.now = moment().format(format)
        m.redraw()
    },
}

const Time = {
    oninit(vnode) {
        state.now = moment().format(format)
        state.timer = setInterval(state.update, 10000)
    },
    onremove(vnode) {
        clearInterval(state.timer)
    },
    view(vnode) {
        return m('li.nav-item', m('span.navbar-text.text-primary.mr-4', state.now))
    }
}
export default Time