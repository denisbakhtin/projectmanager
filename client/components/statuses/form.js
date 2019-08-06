import m from 'mithril'
import error from '../shared/error'
import {
    addSuccess
} from '../shared/notifications'
import state from './state'

export default function form() {
    return [
        m('.form-group', [
            m('label', 'Status name'),
            m('input.form-control[type=text]', {
                oncreate: (el) => {
                    el.dom.focus()
                },
                oninput: (e) => state.setName(e.target.value),
                value: state.status.name
            })
        ]),
        m('.form-group w-25', [
            m('label', 'Order'),
            m('input.form-control[type=number][min=0]', {
                oninput: (e) => state.setOrder(e.target.value),
                value: state.status.order
            })
        ]),
        m('.form-group', [
            m('label', "Description"),
            m('textarea.form-control', {
                oninput: (e) => state.setDescription(e.target.value),
                value: state.status.description
            })
        ]),
        m('.mb-2', m(error, {
            errors: state.errors
        })),
    ]
}