import m from 'mithril'
import error from '../shared/error'
import service from '../../utils/service.js'
import projects_item from './projects_item.js'

const Filters = Object.freeze({
    ALL: (project) => true,
    OPEN: (project) => project.archived == false,
    ARCHIVED: (project) => project.archived == true,
    FAVORITE: (project) => project.favorite == true,
})

export default function Projects() {
    let projects = [],
        errors = [],
        filter,

        activeClass = (fil) => (filter === Filters[fil]) ? "active" : "",
        setFilter = (fil) => {
            localStorage.projectFilter = fil
            filter = Filters[fil]
        },
        getFilter = () => localStorage.projectFilter,

        getAll = () =>
            service.getProjects()
                .then((result) => projects = result.slice(0))
                .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            getAll()
            filter = Filters[getFilter() ?? "ALL"] ?? Filters["ALL"]
        },

        view(vnode) {
            let filteredProjects = projects.filter(filter)

            return m(".projects", [
                m('h1.title', 'Projects'),
                m('.filters', [
                    m('button.btn.btn-link', { class: activeClass("ALL"), onclick: () => setFilter("ALL") }, "All"),
                    m('button.btn.btn-link', { class: activeClass("OPEN"), onclick: () => setFilter("OPEN") }, "Open"),
                    m('button.btn.btn-link', { class: activeClass("ARCHIVED"), onclick: () => setFilter("ARCHIVED") }, "Archived"),
                    m('button.btn.btn-link', { class: activeClass("FAVORITE"), onclick: () => setFilter("FAVORITE") }, "Favorite"),
                ]),
                (filteredProjects && filteredProjects.length > 0) ? m('ul.dashboard-box.box-list',
                    filteredProjects.map((proj) => m(projects_item, { key: proj.id, project: proj, onUpdate: getAll }))
                ) : m('p.text-muted', 'The list is empty'),
                m(error, { errors: errors }),
                m('button#floating-add-button.btn.btn-primary[type=button]', {
                    onclick: () => m.route.set('/projects/new')
                },
                    m('i.fa.fa-plus'),
                ),
            ])
        }
    }
}
