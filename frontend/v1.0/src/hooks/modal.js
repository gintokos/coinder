import { useCallback, useState, useEffect } from "react"

export const useModal = () => {
    const [isOpen, setIsOpen] = useState(false)

    const open = useCallback(() => setIsOpen(true), [])

    const close = useCallback(() => setIsOpen(false), [])

    useEffect(() => {
        const handleEscape = (event) => {
            if (event.key === 'Escape') {
            close();
            }
        };

        if (isOpen) {
            document.addEventListener('keydown', handleEscape);
        }

        return () => {
            document.removeEventListener('keydown', handleEscape);
        };
        }, [isOpen, close])

        return [
            isOpen,
            open,
            close
        ]
}