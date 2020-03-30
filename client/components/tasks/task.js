import m from 'mithril'
import error from '../shared/error'
import MarkdownIt from 'markdown-it'
import service from '../../utils/service.js'
import {
    responseErrors,
    humanTaskSpent
} from '../../utils/helpers'
import comments from '../comments/comments'
import { startTask } from '../shared/active_task'
import files from '../attached_files/files'
import edit_comment_modal from '../comments/edit_comment_modal'

export default function Task() {
    let task,
        id,
        errors = [],
        md,
        showSolutionModal = false,

        get = () =>
            service.getTask(id)
                .then((result) => task = result)
                .catch((error) => errors = responseErrors(error)),

        spent = () => (task.task_logs && task.task_logs.length > 0) ? humanTaskSpent(task, true) : ''

    return {
        oninit(vnode) {
            id = m.route.param('id')
            get()
            md = new MarkdownIt()
        },

        view(vnode) {
            return m(".task", (task) ? [
                m('h1.title', [
                    task.name,
                    (task.category && task.category.id > 0) ?
                        m('a.badge.badge-light.ml-2', { onclick: () => m.route.set('/categories/' + task.category.id) }, [
                            m('i.fa.fa-tag.mr-1'),
                            task.category.name
                        ]) : null,
                ]),
                (spent() != '') ? m('p.task-time-spent', [
                    m('i.fa.fa-clock-o.mr-2'),
                    'Total time spent: ' + spent(),
                ]) : null,
                m('.buttons', [
                    m('button.btn.btn-primary.btn-raised.btn-round[type=button]', {
                        onclick: () => startTask(task, () => get())
                    }, [
                        m('i.fa.fa-play'),
                        'Start',
                    ]),
                    m('button.btn.btn-default.btn-icon[type=button]', {
                        onclick: () => m.route.set('/tasks/edit/' + task.id)
                    }, m('i.fa.fa-edit')),
                    (!task.completed) ?
                        m('button.btn.btn-primary.btn-icon[type=button]', {
                            onclick: () => showSolutionModal = true,
                        }, m('i.fa.fa-check')) : null,
                    m('button.btn.btn-default.btn-icon[type=button]', {
                        onclick: () => remove()
                    }, m('i.fa.fa-trash-o')),
                ]),
                (task.description) ? m('.task-contents', m.trust(md.render(task.description))) : null,
                m(files, { files: task.files, readOnly: true }),
                m(comments, { comments: task.comments, task_id: task.id, onUpdate: () => get() }),

                m(error, { errors: errors }),
                m('.actions.mt-4', [
                    m('button.btn.btn-outline-secondary.mr-2[type=button]', {
                        onclick: () => window.history.back()
                    }, "Back"),
                ]),

                (showSolutionModal) ? m(edit_comment_modal, {
                    task_id: task.id,
                    is_solution: true,
                    onOk: () => { showSolutionModal = false; get(); },
                    onCancel: () => { showSolutionModal = false },
                }) : null,
            ] : m('Loading...'))
        }
    }
}
