import m from 'mithril'
import {
    humanDate,
    responseErrors
} from '../../utils/helpers'
import {
    addDanger
} from '../shared/notifications'
import service from '../../utils/service.js'
import yesno_modal from '../shared/yesno_modal'

export default function PagesItem() {
    let errors = [],
        onUpdate,
        showModal = false,

        remove = (page) =>
            service.deletePage(page.id)
                .then((result) => onUpdate())
                .catch((error) => addDanger(responseErrors(error).join('. ')))

    return {
        oninit(vnode) {
            onUpdate = vnode.attrs.onUpdate ?? (() => null)
        },

        view(vnode) {
            let page = vnode.attrs.page

            return m('li', [
                m('.item-description', [
                    m('h3.item-title', [
                        page.name,
                        (page.published) ? m('span.badge.badge-success.ml-2', 'Published') : null,
                    ]),
                    m('.dates', [
                        m('span.fa.fa-calendar'),
                        m('span', 'Created on: '),
                        m('span', humanDate(page.created_at)),
                        page.updated_at > page.created_at ? [
                            m('span.fa.fa-calendar.ml-3'),
                            m('span', 'Updated on: '),
                            m('span', humanDate(page.updated_at)),
                        ] : null,
                    ]),
                ]),
                m('.buttons', [
                    m('a.btn.btn-primary.btn-raised.btn-round[target=_blank]', {
                        href: "/pages/" + page.id,
                    }, 'Details'),
                    m('button.btn.btn-default.btn-icon[type=button]', {
                        onclick: () => m.route.set('/pages/edit/' + page.id)
                    }, m('i.fa.fa-edit')),
                    m('button.btn.btn-default.btn-icon[type=button]', {
                        onclick: () => showModal = true
                    }, m('i.fa.fa-trash-o')),
                ]),
                (showModal) ? m(yesno_modal, {
                    onYes: () => { remove(page); showModal = false },
                    onNo: () => showModal = false
                }) : null,
            ])
        }
    }
}
