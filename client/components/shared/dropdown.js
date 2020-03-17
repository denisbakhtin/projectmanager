import m from 'mithril'

export default function Dropdown() {
    let show = false,
        key = 'dropdown',
        keepState = false,

        loadState = () => {
            let s = ''
            if (keepState)
                s = sessionStorage.getItem(key)
            show = (s) ? JSON.parse(s) : show
        },
        storeState = (val) => {
            show = val
            if (keepState)
                sessionStorage.setItem(key, JSON.stringify(val))
        }

    return {
        oninit: (vnode) => {
            key = vnode.attrs.id || key
            keepState = vnode.attrs.keepState || keepState
            loadState()
        },
        view: (vnode) => {
            for (let child of vnode.attrs.children) {
                //prevent default navigation to /#
                if (child.tag == "a")
                    child.attrs['onclick'] = (e) => e.preventDefault()
                //prevent event propagation to parent item which leads to dropdown collapse
                if (child.tag == "div" && child.children.length > 0)
                    for (let item of child.children) {
                        if (item) item.attrs['onclick'] = (e) => e.stopPropagation()
                    }
            }

            return m('li.nav-item.dropdown', {
                class: show ? 'show' : '',
                onclick: (e) => storeState(!show),
            }, [
                m('div.dropdown-overlay'),
                vnode.attrs.children
            ])
        }
    }
}
