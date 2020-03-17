import m from 'mithril'
import {
    humanDate,
    responseErrors
} from '../../utils/helpers'
import {
    addDanger,
    addSuccess
} from '../shared/notifications'
import service from '../../utils/service.js'
import yesno_modal from '../shared/yesno_modal'

export default function ProjectsItem() {
    let onUpdate,
        showModal = false,
        toggleArchive = (project) => {
            project.archived = !project.archived
            service.archiveProject(project.id, project).
                then((result) => {
                    if (project.archived)
                        addSuccess("Project archived")
                    else
                        addSuccess("Project unarchived")
                }).catch((error) => addDanger(responseErrors(error).join('. ')))
        },
        toggleFavor = (project) => {
            project.favorite = !project.favorite
            service.favorProject(project.id, project).
                then((result) => {
                    if (project.favorite)
                        addSuccess("Project is now favorite")
                    else
                        addSuccess("Project is favorite no more")
                }).catch((error) => addDanger(responseErrors(error).join('. ')))
        },

        remove = (project) =>
            service.deleteProject(project.id)
                .then((result) => onUpdate())
                .catch((error) => addDanger(responseErrors(error).join('. ')))

    return {
        oninit(vnode) {
            onUpdate = vnode.attrs.onUpdate ?? (() => null)
        },

        view(vnode) {
            let project = vnode.attrs.project

            return m('li', [
                m('.item-description', [
                    m('h3.item-title', [
                        project.name,
                        (project.category.id > 0) ?
                            m('a.badge.badge-light.ml-2', { onclick: () => m.route.set('/categories/' + project.category.id) }, [
                                m('i.fa.fa-tag.mr-1'),
                                project.category.name
                            ]) : null,
                        (!project.archived) ? m('span.badge.badge-success.ml-2', 'Open') : null,
                    ]),
                    m('.dates', [
                        m('span.fa.fa-calendar'),
                        m('span', 'Created on: '),
                        m('span', humanDate(project.created_at)),
                        project.updated_at > project.created_at ? [
                            m('span.fa.fa-calendar.ml-3'),
                            m('span', 'Updated on: '),
                            m('span', humanDate(project.updated_at)),
                        ] : null,
                    ]),
                ]),
                m('.buttons', [
                    m('button.btn.btn-primary.btn-raised.btn-round[type=button]', {
                        onclick: () => m.route.set('/projects/' + project.id)
                    }, [
                        'Details',
                        (project.tasks && project.tasks.length > 0) ? m('span.badge.badge-primary.ml-2', project.tasks.length) : '',
                    ]),
                    m('button.btn.btn-default.btn-icon[type=button]', {
                        title: "Edit",
                        onclick: () => m.route.set('/projects/edit/' + project.id)
                    }, m('i.fa.fa-edit')),
                    m('button.btn.btn-default.btn-icon[type=button]', {
                        title: (project.favorite) ? "Remove from favorites" : "Move to favorites",
                        onclick: () => toggleFavor(project),
                    }, (project.favorite) ? m('i.fa.fa-star') : m('i.fa.fa-star-o')),
                    m('button.btn.btn-default.btn-icon[type=button]', {
                        title: (project.archived) ? "Unarchive" : "Archive",
                        onclick: () => toggleArchive(project),
                    }, m('i.fa.fa-archive')),
                    m('button.btn.btn-default.btn-icon[type=button]', {
                        title: "Delete",
                        onclick: () => showModal = true
                    }, m('i.fa.fa-trash-o')),
                ]),
                (showModal) ? m(yesno_modal, {
                    onYes: () => { remove(project); showModal = false },
                    onNo: () => showModal = false
                }) : null,
            ])
        }
    }
}
