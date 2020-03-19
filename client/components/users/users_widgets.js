import m from 'mithril'
import service from '../../utils/service.js'
import { responseErrors } from '../../utils/helpers'

export function UsersCountWidget() {
    let count = 0,
        errors,

        get = () =>
            service.getUsersSummary()
                .then((result) => count = result.count)
                .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            get()
        },

        view(vnode) {

            return m(".card.count-widget",
                m('a.card-body[href=#!/users]', [
                    m('.count', count),
                    m('.description', 'Users'),
                    (errors) ? m('i.fa.fa-exclamation-circle.error-icon', { title: responseErrors(errors) }) : null,
                ])
            )
        }
    }
}
