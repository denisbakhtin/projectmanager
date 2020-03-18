import m from 'mithril'
import {
    humanDate
} from '../../utils/helpers'
import error from '../shared/error'
import service from '../../utils/service.js'
import categories_item from './categories_item'

export default function Categories() {
    let categories = [],
        errors = [],

        getAll = () =>
            service.getCategories()
                .then((result) => categories = result.slice(0))
                .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            getAll()
        },

        view(vnode) {
            return m(".categories", [
                m('h1.title.mb-4', 'Categories of Projects & Tasks'),
                categories.length > 0 ?
                    m('ul.dashboard-box.box-list', [
                        categories.map((cat) => m(categories_item, { key: cat.id, category: cat, onUpdate: getAll }))
                    ]) : m('p.text-muted', 'No categories yet.'),
                m(error, { errors: errors }),
                m('.actions.mt-4', [
                    m('button.btn.btn-primary[type=button]', {
                        onclick: () => m.route.set('/categories/new')
                    }, [
                        m('i.fa.fa-plus.mr-1'),
                        "New category"
                    ])
                ]),
            ])
        }
    }
}
