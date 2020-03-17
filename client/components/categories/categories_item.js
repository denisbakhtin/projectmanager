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

export default function CategoriesItem() {
    let errors = [],
        onUpdate,
        showModal = false,

        remove = (cat) =>
            service.deleteCategory(cat.id)
                .then((result) => onUpdate())
                .catch((error) => addDanger(responseErrors(error).join('. ')))

    return {
        oninit(vnode) {
            onUpdate = vnode.attrs.onUpdate ?? (() => null)
        },

        view(vnode) {
            let category = vnode.attrs.category

            return m('li', [
                m('.item-description', [
                    m('h3.item-title', [
                        category.name,
                    ]),
                    m('.dates', [
                        m('span.fa.fa-calendar'),
                        m('span', 'Created on: '),
                        m('span', humanDate(category.created_at)),
                        category.updated_at > category.created_at ? [
                            m('span.fa.fa-calendar.ml-3'),
                            m('span', 'Updated on: '),
                            m('span', humanDate(category.updated_at)),
                        ] : null,
                    ]),
                ]),
                m('.buttons', [
                    m('button.btn.btn-primary.btn-raised.btn-round[type=button]', {
                        onclick: () => m.route.set('/categories/' + category.id)
                    }, 'Details'),
                    m('button.btn.btn-default.btn-icon[type=button]', {
                        onclick: () => m.route.set('/categories/edit/' + category.id)
                    }, m('i.fa.fa-edit')),
                    m('button.btn.btn-default.btn-icon[type=button]', {
                        onclick: () => showModal = true
                    }, m('i.fa.fa-trash-o')),
                ]),
                (showModal) ? m(yesno_modal, {
                    onYes: () => { remove(category); showModal = false },
                    onNo: () => showModal = false
                }) : null,
            ])
        }
    }
}
