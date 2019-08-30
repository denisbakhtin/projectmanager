import m from 'mithril'
import dropdown from './dropdown'
import state from './state'
import {
	entityUrl
} from '../../utils/helpers'

function Notifications(initialVnode) {
	return {
		oninit: () => state.getNotifications(),
		view: (vnode) => {
			return m(dropdown, {
				children: [
					m('a.nav-link#navbar-notifications[href=#]', [
						m('span.fa.fa-bell-o'),
						(state.notifications.length > 0) ? m('span.badge.badge-pill.badge-primary', state.notifications.length) : false,
					]),
					(state.notifications.length > 0) ? m('.dropdown-menu', state.notifications.map((n) => m('a.dropdown-item', {
						href: entityUrl(n.entity, n.entity_id)
					}, n.title))) : false,
				]
			})
		}
	}
}
export default Notifications;