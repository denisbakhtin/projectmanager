import m from 'mithril'
import MarkdownIt from 'markdown-it'
import error from '../shared/error'
import service from '../../utils/service.js'
import {
    responseErrors,
} from '../../utils/helpers'
import tasks_item from '../tasks/tasks_item.js'
import files from '../attached_files/files'

export default function Project() {
    let project = {},
        errors = [],
        id,
        md,

        get = () =>
            service.getProject(id)
                .then((result) => project = result)
                .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            id = m.route.param('id')
            get()
            md = new MarkdownIt()
        },

        view(vnode) {
            return m(".project", (project) ? [
                m('h1.title', [
                    project.name,
                    (project.category && project.category.id > 0) ?
                        m('a.badge.badge-light.ml-2', { onclick: () => m.route.set('/categories/' + project.category.id) }, [
                            m('i.fa.fa-tag.mr-1'),
                            project.category.name
                        ]) : null,
                    m('button.ml-2.btn.btn-sm.btn-default[type=button]', {
                        onclick: () => m.route.set('/projects/edit/' + project.id)
                    }, 'Edit'),
                ]),
                m('p.project-time-spent', [
                    m('i.fa.fa-clock-o.mr-2'),
                    'Total time spent: 17 minutes',
                ]),
                (project.description) ? m('p.project-contents', m.trust(md.render(project.description))) : null,
                m(files, { files: project.files, readOnly: true }),
                (project.tasks && project.tasks.length > 0) ?
                    m('ul.dashboard-box.box-list',
                        project.tasks.map((task) => m(tasks_item, { task: task, onUpdate: get }))
                    ) : null,
                m(error, { errors: errors }),
                m('.actions.mt-4', [
                    m('button.btn.btn-primary[type=button]', {
                        onclick: () => m.route.set('/tasks/new?project_id=' + project.id)
                    }, [
                        m('i.fa.fa-plus.mr-1'),
                        "New task"
                    ]),
                    m('button.btn.btn-outline-secondary.mr-2[type=button]', {
                        onclick: () => window.history.back()
                    }, "Back"),
                ])
            ] : m('Loading...'))
        }
    }
}
