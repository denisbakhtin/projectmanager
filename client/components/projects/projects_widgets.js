import m from 'mithril'
import service from '../../utils/service.js'
import { responseErrors } from '../../utils/helpers'

export function ProjectsCountWidget() {
    let count = 0,
        errors,

        get = () =>
            service.getProjectsSummary()
                .then((result) => count = result.count)
                .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            get()
        },

        view(vnode) {

            return m(".card.count-widget",
                m('a.card-body[href=#!/projects]', [
                    m('.count', count),
                    m('.description', 'Projects'),
                    (errors) ? m('i.fa.fa-exclamation-circle.error-icon', { title: responseErrors(errors) }) : null,
                ])
            )
        }
    }
}
