import m from 'mithril'
import {
    humanDate,
    humanTaskSpent,
    responseErrors,
    isZeroDate
} from '../../utils/helpers'
import service from '../../utils/service.js'
import edit_comment_modal from '../comments/edit_comment_modal'
import { startTask } from '../shared/active_task'
import { addDanger } from '../shared/notifications'
import yesno_modal from '../shared/yesno_modal'
import button_menu from '../shared/button_menu'

export default function TasksItem() {
    let onUpdate,
        onOpenClick,
        onExpiredClick,
        showCommentsModal = false,
        showRemoveModal = false,
        isSolution = false,

        remove = (task) =>
            service.deleteTask(task.id)
                .then((result) => onUpdate())
                .catch((error) => addDanger(responseErrors(error).join('. '))),

        spent = (task) => (task.task_logs && task.task_logs.length > 0) ? humanTaskSpent(task, true) : ''

    return {
        oninit(vnode) {
            onUpdate = vnode.attrs.onUpdate ?? (() => null)
            onOpenClick = vnode.attrs.onOpenClick ?? (() => null)
            onExpiredClick = vnode.attrs.onExpiredClick ?? (() => null)
        },

        view(vnode) {
            let task = vnode.attrs.task

            return m('li', { class: 'priority' + task.priority }, [
                m('.item-description', [
                    m('h3.item-title', [
                        m('span.mr-2', task.name),
                        (task.category.id > 0) ?
                            m('a.badge.badge-light.badge-category.mr-2', { onclick: () => m.route.set('/categories/' + task.category.id) }, [
                                m('i.fa.fa-tag.mr-1'),
                                task.category.name
                            ]) : null,
                        (!task.completed) ? m('a.badge.badge-success', { onclick: onOpenClick }, "Open") : null,
                        (!!task.end_date && Date.parse(task.end_date) < Date.now()) ? m('a.badge.badge-warning', { onclick: onExpiredClick }, "Expired") : null,
                    ]),
                    (!isZeroDate(task.start_date) || !isZeroDate(task.end_date) || spent(task) != '') ? m('.dates', [
                        (!isZeroDate(task.start_date)) ? m('span.created-on.mr-3', [
                            m('span.fa.fa-calendar'),
                            m('span', 'Starts: '),
                            m('span', humanDate(task.start_date)),
                        ]) : null,
                        (!isZeroDate(task.end_date)) ? m('span.updated-on.mr-3', [
                            m('span.fa.fa-calendar'),
                            m('span', 'Ends: '),
                            m('span', humanDate(task.end_date)),
                        ]) : null,
                        (spent(task) != '') ? m('span.time-spent', { title: "Total time spent" }, [
                            m('span.fa.fa-clock-o'),
                            spent(task),
                        ]) : null,
                    ]) : null,
                ]),
                m('.buttons', [
                    m('button.btn.btn-primary.btn-raised.btn-round[type=button]', {
                        onclick: () => startTask(task, onUpdate)
                    }, [
                        m('i.fa.fa-play'),
                        m('span.button-text', 'Start'),
                    ]),
                    m('button.btn.btn-default.btn-round[type=button]', {
                        onclick: () => m.route.set('/tasks/' + task.id)
                    }, [
                        m('i.fa.fa-info'),
                        m('span.button-text.mr-2', 'Details'),
                    ]),
                    m(button_menu, {
                        children: [
                            m('.dropdown-menu', [
                                m('button.dropdown-item.btn.btn-default.btn-icon[type=button]', {
                                    onclick: () => m.route.set('/tasks/edit/' + task.id)
                                }, [
                                    m('i.fa.fa-edit'),
                                    m('span.text.ml-1', 'Edit')
                                ]),
                                m('button.dropdown-item.btn.btn-default.btn-icon[type=button]', {
                                    onclick: () => { isSolution = false; showCommentsModal = true }
                                }, [
                                    m('i.fa.fa-commenting-o'),
                                    m('span.text.ml-1', 'Comment')
                                ]),
                                (!task.completed) ?
                                    m('button.dropdown-item.btn.btn-primary.btn-icon[type=button]', {
                                        onclick: () => { isSolution = true; showCommentsModal = true }
                                    }, [
                                        m('i.fa.fa-check'),
                                        m('span.text.ml-1', 'Solve')
                                    ]) : null,
                                m('button.dropdown-item.btn.btn-default.btn-icon[type=button]', {
                                    onclick: () => showRemoveModal = true
                                }, [
                                    m('i.fa.fa-trash-o'),
                                    m('span.text.ml-1', 'Delete')
                                ]),
                            ])
                        ]
                    })
                ]),

                (showCommentsModal) ? m(edit_comment_modal, {
                    task_id: task.id,
                    is_solution: isSolution,
                    onOk: () => { showCommentsModal = false; onUpdate(); },
                    onCancel: () => { showCommentsModal = false },
                }) : null,

                (showRemoveModal) ? m(yesno_modal, {
                    onYes: () => { remove(task); showRemoveModal = false },
                    onNo: () => showRemoveModal = false
                }) : null,
            ])
        }
    }
}
