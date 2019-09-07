import m from 'mithril'
import moment from 'moment'

const format = 'MMMM Do YYYY, h:mm a'

export default function Time() {
    let now = '',
        timer = null,

        update = () => {
            now = moment().format(format)
            m.redraw()
        }

    return {
        oninit(vnode) {
            now = moment().format(format)
            timer = setInterval(update, 10000)
        },
        onremove(vnode) {
            clearInterval(timer)
        },
        view(vnode) {
            return m('li.nav-item', m('span.navbar-text.text-primary.mr-4', now))
        }
    }
}
