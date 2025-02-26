import { useCallback, useState, useEffect } from "react"

export const useModal = () => {
    const [isOpen, setIsOpen] = useState(false)

    const open = useCallback(() => {
        setIsOpen(true)
        window.history.pushState(null, '', window.location.href)
    }, [])

    const close = useCallback(() => {
        setIsOpen(false)
    }, [])

    useEffect(() => {
        const handleEscape = (event) => {
            if (event.key === 'Escape') {
                close()
            }
        }

        const handlePopState = () => {
            if (isOpen) {
                close()
            }
        }

        if (isOpen) {
            document.addEventListener('keydown', handleEscape)
            window.addEventListener('popstate', handlePopState)
        }

        return () => {
            document.removeEventListener('keydown', handleEscape)
            window.removeEventListener('popstate', handlePopState)
        }
    }, [isOpen, close])

    return [
        isOpen,
        open,
        close
    ]
}