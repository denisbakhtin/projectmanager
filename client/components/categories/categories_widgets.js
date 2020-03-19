import m from 'mithril'
import service from '../../utils/service.js'
import { responseErrors } from '../../utils/helpers'

export function CategoriesCountWidget() {
    let count = 0,
        errors,

        get = () =>
            service.getCategoriesSummary()
                .then((result) => count = result.count)
                .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            get()
        },

        view(vnode) {

            return m(".card.count-widget",
                m('a.card-body[href=#!/categories]', [
                    m('.count', count),
                    m('.description', 'Categories'),
                    (errors) ? m('i.fa.fa-exclamation-circle.error-icon', { title: responseErrors(errors) }) : null,
                ])
            )
        }
    }
}
