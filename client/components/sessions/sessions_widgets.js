import m from 'mithril'
import service from '../../utils/service.js'
import { responseErrors } from '../../utils/helpers'

export function SessionsCountWidget() {
    let count = 0,
        errors,

        get = () =>
            service.getSessionsSummary()
                .then((result) => count = result.count)
                .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            get()
        },

        view(vnode) {

            return m(".card.count-widget",
                m('a.card-body[href=#!/sessions]', [
                    m('.count', count),
                    m('.description', 'Sessions'),
                    (errors) ? m('i.fa.fa-exclamation-circle.error-icon', { title: responseErrors(errors) }) : null,
                ])
            )
        }
    }
}
