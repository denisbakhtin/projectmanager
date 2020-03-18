import m from 'mithril'
import error from '../shared/error'
import service from '../../utils/service.js'
import pages_item from './pages_item'

export default function Pages() {
    let pages = [],
        errors = [],

        getAll = () =>
            service.getPages()
                .then((result) => pages = result.slice(0))
                .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            getAll()
        },

        view(vnode) {
            return m(".pages", [
                m('h1.title.mb-4', 'Public Pages'),
                pages.length > 0 ?
                    m('ul.dashboard-box.box-list', [
                        pages.map((page) => m(pages_item, { key: page.id, page: page, onUpdate: getAll }))
                    ]) : m('p.text-muted', 'No pages yet.'),
                m(error, { errors: errors }),
                m('.actions.mt-4', [
                    m('button.btn.btn-primary[type=button]', {
                        onclick: () => m.route.set('/pages/new')
                    }, [
                        m('i.fa.fa-plus.mr-1'),
                        "New page"
                    ])
                ]),
            ])
        }
    }
}
