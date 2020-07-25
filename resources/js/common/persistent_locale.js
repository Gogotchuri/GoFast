export const getLocale = () => {
    let locale = window.localStorage.getItem("locale");
    if (!locale) return "en";
    return locale;
};

export const changeLocale = locale => {
    window.localStorage.setItem("locale", locale);
};