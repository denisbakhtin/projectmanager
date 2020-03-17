import './css/site.scss'
import m from 'mithril'
import 'moment'
import users from './components/users/users'
import login from './components/account/login'
import logout from './components/account/logout'
//import roles from './components/roles/roles'
import register from './components/account/register'
import activationNotice from './components/account/activation_notice'
import activate from './components/account/activate'
import reset from './components/account/reset'
import resetNotice from './components/account/reset_notice'
import manage from './components/account/manage'
import home from './components/home'
import under_construction from './components/shared/under_construction'

import categories from './components/categories/categories'
import edit_category from './components/categories/edit_category'
import category from './components/categories/category'

import projects from './components/projects/projects'
import edit_project from './components/projects/edit_project'
import project from './components/projects/project'

import settings from './components/settings/settings'
import edit_setting from './components/settings/edit_setting.js'

import tasks from './components/tasks/tasks'
import edit_task from './components/tasks/edit_task'
import task from './components/tasks/task'

import sessions from './components/sessions/sessions'
import edit_session from './components/sessions/edit_session'
import session from './components/sessions/session'

import spent_report from './components/reports/spent'

import layout from './components/shared/layout'
import public_layout from './components/shared/layout_public'
import Auth from './utils/auth'

import search from './components/search/search'

import pages from './components/pages/pages'
import edit_page from './components/pages/edit_page'

const app_root = document.getElementById("app-root");

//render component with layout
function withLayout(comp) {
    return {
        view: () => Auth.isLoggedIn() ? m(layout, {
            child: comp
        }) : m(public_layout, {
            child: comp
        })
    }
}

//Authorization route filter
function checkAuthorized(comp) {
    return function (args, path) {
        if (!Auth.isLoggedIn()) {
            localStorage.returnURL = path
            Auth.logout()
        } else return withLayout(comp)
    }
}

//Authorization route filter for admins
function checkAuthorizedAdmin(comp) {
    return function (args, path) {
        if (!Auth.isLoggedIn() || !Auth.isAdmin()) {
            localStorage.returnURL = path
            Auth.logout()
        } else return withLayout(comp)
    }
}

//route wrapper
function route(comp, requiresAuth = true, requiresAdmin = false) {
    return (requiresAuth) ? (requiresAdmin) ? {
        //admin user
        onmatch: checkAuthorizedAdmin(comp)
    } : {
            //ordinary user    
            onmatch: checkAuthorized(comp)
        } : withLayout(comp) //not authenticated
}

const routes = {
    '/': route(home, false),
    '/login': route(login, false),
    '/register': route(register, false),
    '/activation_notice': route(activationNotice, false),
    '/activate/:token': route(activate, false),
    '/reset': route(reset, false),
    '/reset/:token': route(reset, false),
    '/reset_notice': route(resetNotice, false),
    '/manage': route(manage),
    //'/roles': route(roles),
    '/users': route(users, true, true),

    '/categories': route(categories),
    '/categories/new': route(edit_category),
    '/categories/edit/:id': route(edit_category),
    '/categories/:id': route(category),

    '/reports/spent': route(spent_report),

    '/sessions': route(sessions),
    '/sessions/new': route(edit_session),
    '/sessions/:id': route(session),

    '/projects': route(projects),
    '/projects/new': route(edit_project),
    '/projects/edit/:id': route(edit_project),
    '/projects/:id': route(project),

    '/tasks': route(tasks),
    '/tasks/new': route(edit_task),
    '/tasks/edit/:id': route(edit_task),
    '/tasks/:id': route(task),

    '/settings': route(settings, true, true),
    '/settings/new': route(edit_setting, true, true),
    '/settings/edit/:id': route(edit_setting, true, true),

    '/pages': route(pages, true, true),
    '/pages/new': route(edit_page, true, true),
    '/pages/edit/:id': route(edit_page, true, true),

    '/support': route(under_construction),
    '/lock': route(under_construction),
    '/download': route(under_construction),
    '/web_site': route(under_construction),
    '/phone': route(under_construction),

    '/search': route(search),

    '/logout': route(logout, false),
};



m.route(app_root, "/", routes)

//turn on webpack hot reload
if (module.hot) {
    module.hot.accept()
}
