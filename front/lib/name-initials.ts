const getNameInitials = (nome: string): string => {
    const palavras = nome.split(' ');
    if(palavras.length === 1)
        return palavras[0].slice(0, 2).toUpperCase()

    return `${palavras[0][0]}${palavras[palavras.length - 1][0]}`.toUpperCase()
}

export default getNameInitials