import m from 'mithril'
import {
    humanDate,
    humanTaskSpent,
    responseErrors
} from '../../utils/helpers'
import service from '../../utils/service.js'
import edit_comment_modal from '../comments/edit_comment_modal'
import { startTask } from '../shared/active_task'
import { addDanger } from '../shared/notifications'
import yesno_modal from '../shared/yesno_modal'

export default function TasksItem() {
    let onUpdate,
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
        },

        view(vnode) {
            let task = vnode.attrs.task

            return m('li', { class: 'priority' + task.priority }, [
                m('.item-description', [
                    m('h3.item-title', [
                        task.name,
                        (task.category.id > 0) ?
                            m('a.badge.badge-light.ml-2', { onclick: () => m.route.set('/categories/' + task.category.id) }, [
                                m('i.fa.fa-tag.mr-1'),
                                task.category.name
                            ]) : null,
                        (!task.completed) ? m('a.badge.badge-success.ml-2', "Open") : null,
                    ]),
                    m('.dates', [
                        m('span.fa.fa-calendar'),
                        m('span', 'Created on: '),
                        m('span', humanDate(task.created_at)),
                        task.updated_at > task.created_at ? [
                            m('span.fa.fa-calendar.ml-3'),
                            m('span', 'Updated on: '),
                            m('span', humanDate(task.updated_at)),
                        ] : null,
                        (spent(task) != '') ? m('span.time-spent.ml-3', { title: "Total time spent" }, [
                            m('span.fa.fa-clock-o'),
                            spent(task),
                        ]) : null,
                    ]),
                ]),
                m('.buttons', [
                    m('button.btn.btn-primary.btn-raised.btn-round[type=button]', {
                        onclick: () => startTask(task, onUpdate)
                    }, [
                        m('i.fa.fa-play'),
                        'Start',
                    ]),
                    m('button.btn.btn-default.btn-round[type=button]', {
                        onclick: () => m.route.set('/tasks/' + task.id)
                    }, [
                        'Details',
                        (task.comments && task.comments.length > 0) ? m('span.comments-count.ml-2.text-muted', [
                            m('i.fa.fa-comments'),
                            task.comments.length
                        ]) : null,
                    ]),
                    m('button.btn.btn-default.btn-icon[type=button]', {
                        onclick: () => m.route.set('/tasks/edit/' + task.id)
                    }, m('i.fa.fa-edit')),
                    m('button.btn.btn-default.btn-icon[type=button]', {
                        onclick: () => { isSolution = false; showCommentsModal = true }
                    }, m('i.fa.fa-commenting-o')),
                    (!task.completed) ?
                        m('button.btn.btn-primary.btn-icon[type=button]', {
                            onclick: () => { isSolution = true; showCommentsModal = true }
                        }, m('i.fa.fa-check')) : null,
                    m('button.btn.btn-default.btn-icon[type=button]', {
                        onclick: () => showRemoveModal = true
                    }, m('i.fa.fa-trash-o')),
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
