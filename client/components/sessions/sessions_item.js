import m from 'mithril'
import {
    humanDate,
    humanSessionSpent,
    responseErrors
} from '../../utils/helpers'
import service from '../../utils/service.js'
import { addDanger } from '../shared/notifications'
import yesno_modal from '../shared/yesno_modal'

export default function SessionsItem() {
    let onUpdate,
        showModal = false,

        remove = (session) =>
            service.deleteSession(session.id)
                .then((result) => onUpdate())
                .catch((error) => addDanger(responseErrors(error).join('. ')))

    return {
        oninit(vnode) {
            onUpdate = vnode.attrs.onUpdate ?? (() => null)
        },

        view(vnode) {
            let session = vnode.attrs.session

            return m('li', [
                m('.item-description', [
                    m('h3.item-title', [
                        'Session #' + session.id,
                    ]),
                    (session.contents && session.contents.length > 0) ? m('p', session.contents) : null,
                    m('.dates', [
                        m('span.fa.fa-calendar'),
                        m('span', 'Created on: '),
                        m('span', humanDate(session.created_at)),
                        (humanSessionSpent(session) != '') ? m('span.time-spent.ml-3', { title: "Total time spent" }, [
                            m('span.fa.fa-clock-o'),
                            humanSessionSpent(session),
                        ]) : null,
                    ]),
                ]),
                m('.buttons', [
                    m('button.btn.btn-primary.btn-raised.btn-round[type=button]', {
                        onclick: () => m.route.set('/sessions/' + session.id)
                    }, 'Details'),
                    m('button.btn.btn-default.btn-icon[type=button]', {
                        onclick: () => showModal = true
                    }, m('i.fa.fa-trash-o')),
                ]),
                (showModal) ? m(yesno_modal, {
                    onYes: () => { remove(session); showModal = false },
                    onNo: () => showModal = false
                }) : null,
            ])
        }
    }
}
