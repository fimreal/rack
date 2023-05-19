# Fimreal's bashrc
# reference: https://wiki.archlinux.org/title/User:Erus_Iluvatar/bashrc
if [[ $- != *i* ]]; then
    # Shell is non-interactive.  Be done now!
    return
fi

# Bash won't get SIGWINCH if another process is in the foreground.
# Enable checkwinsize so that bash will check the terminal size when
# it regains control.  #65623
# http://cnswww.cns.cwru.edu/~chet/bash/FAQ (E11)
shopt -s checkwinsize

# Disable completion when the input buffer is empty.  i.e. Hitting tab
# and waiting a long time for bash to expand all of $PATH.
shopt -s no_empty_cmd_completion

# Enable history appending instead of overwriting when exiting.  #139609
shopt -s histappend

# PS1 color
use_color=false
if type -P dircolors >/dev/null; then
    # Enable colors for ls, etc.  Prefer ~/.dir_colors #64489
    LS_COLORS=
    if [[ -f ~/.dir_colors ]]; then
        eval "$(dircolors -b ~/.dir_colors)"
    elif [[ -f /etc/DIR_COLORS ]]; then
        eval "$(dircolors -b /etc/DIR_COLORS)"
    else
        eval "$(dircolors -b)"
    fi
    # Note: We always evaluate the LS_COLORS setting even when it's the
    # default.  If it isn't set, then `ls` will only colorize by default
    # based on file attributes and ignore extensions (even the compiled
    # in defaults of dircolors). #583814
    if [[ -n ${LS_COLORS:+set} ]]; then
        use_color=true
    else
        # Delete it if it's empty as it's useless in that case.
        unset LS_COLORS
    fi
else
    # Some systems (e.g. BSD & embedded) don't typically come with
    # dircolors so we need to hardcode some terminals in here.
    case ${TERM} in
    [aEkx]term* | rxvt* | gnome* | konsole* | screen | tmux | cons25 | *color) use_color=true ;;
    esac
fi

# Move the test for changing the window title in X terminals inside a function
# to be able to reset $PS1 properly inside $PROMPT_COMMAND
change_window_title() {
    case ${TERM} in
    [aEkx]term* | rxvt* | gnome* | konsole* | interix | tmux*)
        PS1='\[\033]0;\u@\h:\w\007\]'
        ;;
    screen*)
        PS1='\[\033k\u@\h:\w\033\\\]'
        ;;
    *)
        unset PS1
        ;;
    esac
}

if ${use_color}; then
    # vscode need tiny shell
    if [[ ${TERM_PROGRAM} == "vscode" ]]; then
        if [[ ${EUID} == 0 ]]; then
            PS1='\[\033[01;34m\]\W \#\[\033[00m\] '
        else
            PS1='\[\033[01;34m\]\W \$\[\033[00m\] '
        fi
    fi

    if [[ ${EUID} == 0 ]]; then
        # Create a function to be used with $PROMPT_COMMAND
        root() {
            # Find the return value of the last command and
            # save it to another variable before anything else,
            # since we are executing other commands after the test
            # but before setting $PS1
            RET=$?
            change_window_title
            if [ $RET -gt 0 ]; then
                PS1+="\[\033[01;35m\]$RET \[\033[01;31m\]\h \[\033[01;34m\]\w \[\033[00m\]\$ "
            else
                PS1+="\[\033[01;31m\]\h \[\033[01;34m\]\w \[\033[00m\]\$ "
            fi
            unset RET
        }
        PROMPT_COMMAND=root
    else
        # Create a function to be used with $PROMPT_COMMAND
        user() {
            # Find the return value of the last command and
            # save it to another variable before anything else,
            # since we are executing other commands after the test
            # but before setting $PS1
            RET=$?
            change_window_title
            if [ $RET -gt 0 ] && [ $RET -ne 130 ]; then
                PS1+="\[\033[01;35m\]$RET \[\033[01;32m\]\u\[\033[00m\]@\h \[\033[01;34m\]\w \[\033[00m\]\$ "
            else
                PS1+="\[\033[01;32m\]\u\[\033[00m\]@\h \[\033[01;34m\]\w \[\033[00m\]\$ "
            fi
            unset RET
        }
        PROMPT_COMMAND=user
    fi

    alias diff='diff --color=auto'
    alias egrep='grep -E --color=auto'
    alias grep='grep --color=auto'
    alias fgrep='grep -F --color=auto'
    alias ip='ip --color=auto'
    alias ls='ls --color=auto --time-style=long-iso'
else
    # show root@ when we don't have colors
    change_window_title
    PS1+='\u@\h \w \$ '
fi

# tiny PS1
function tinyps1() {
    if [[ ${EUID} == 0 ]] ; then
        PS1='\[\033[01;34m\]\W \#\[\033[00m\] '
    else
        PS1='\[\033[01;34m\]\W \$\[\033[00m\] '
    fi
}
if [[ ${TERM_PROGRAM} == "vscode" ]] ; then
    tinyps1
fi

# Try to keep environment pollution down, EPA loves us.
unset use_color


##========== ENV ==========##
# iterm2
export BASH_SILENCE_DEPRECATION_WARNING=1
# timezone
export TZ='Asia/Shanghai'
# HISTORY conf
export HISTSIZE=100000
export HISTFILESIZE=100000
export HISTTIMEFORMAT="%F %T "
export HISTCONTROL=ignoredups
# PATH
export PATH=$PATH:/root/scripts:/root/bin
# EDITOR
if command -v vim &>/dev/null; then
    export EDITOR=vim
    export KUBE_EDITOR="vim"
elif command -v vi &>/dev/null; then
    export EDITOR=vi
fi

# goproxy
export HUGO_MODULE_PROXY=https://goproxy.cn/,direct
export GOPROXY=https://goproxy.cn,direct \
       GO111MOUDULE=on \
     # GOPATH=~/go:~/OneDrive/Code_root/golang:$GOPATH

##========== alias ==========##
alias la='ls -AF'
alias ll='ls -lhF'
alias lla='ls -lAhF'
alias pgrep="fgrep -lf"
alias grep#='grep -v "^#"'
alias grep##='egrep -v "^#|^$"'
alias pyweb="python3 -m http.server"
# alias toolbox='toolbox --bind /server:/server --bind $PWD:/root/current_path/'
alias lzd="docker run -it --rm -v /var/run/docker.sock:/var/run/docker.sock lazyteam/lazydocker"
# K8STOKEN=$(kubectl -n kube-system describe secret default| awk '$1=="token:"{print $2}') \
alias kgp='kubectl get pod --sort-by=".metadata.creationTimestamp"'
alias kga='kubectl get pods --include-uninitialized' # 包括未初始化的
alias kge='kubectl get pod --sort-by=".metadata.creationTimestamp" --field-selector=status.phase!=Running'
alias kgr='kubectl get pods --sort-by=".status.containerStatuses[0].restartCount"'
alias kgi='kubectl get pods  -o jsonpath='"'"'{range .items[*]}{@.metadata.name}{" "}{@.spec.containers[*].image}{"\n"}{end}'"'"''
alias ktk='kubectl -n kube-system describe secret default| awk '"'"'$1=="token:"{print $2}'"'"''
alias ktn='kubectl top node'
# alias kdn='kubectl get pod -o wide --all-namespaces | awk '{if($4!="Running"){cmd="kubectl -n "$1" delete pod "$2; system(cmd)}}''
# alias nginx='kubectl exec $(kubectl get pods -o jsonpath='"'"'{range .items[*]}{@.metadata.name}{"\n"}{end}'"'"' -l app=openresty) -- nginx '
alias gitlog="git log --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit"
alias gomodoff="go env -w GO111MODULE=off"
alias gomodon="go env -w GO111MODULE=on"
alias jet="step crypto jwt inspect --insecure <<<"

if command -v podman &>/dev/null; then
    alias docker='echo '\''podman: hey, Don not use docker, '\'''
fi

##========== completion ==========##
for cliname in kubeadm kubectl helm argocd flux; do
    if command -v $cliname &>/dev/null; then
        source <($cliname completion bash)
    fi
done
if [ -f ~/.git-completion.bash ]; then
    . ~/.git-completion.bash
fi
if command -v brew &>/dev/null; then
    if [ -f $(brew --prefix)/etc/bash_completion ]; then
        . $(brew --prefix)/etc/bash_completion
    fi
fi

##========== functions ==========##
# Terminal proxy
export _PROXY_PS1_ORI=$PS1
function proxy_on() {
    port=$1
    export no_proxy="localhost,127.0.0.1,localaddress,.localdomaiin.com"
    export http_proxy="http://127.0.0.1:${port:-1080}"
    export https_proxy=$http_proxy
    export PS1="\[\e[36;1m\](proxy on)\[\e[35;0m\]$_PROXY_PS1_ORI"
    echo "proxy enabled now. $http_proxy $(timeout 3 curl -s epurs.com/ipinfo | jq -r '. | .query + " " + .country')"
}
function proxy_off() {
    unset http_proxy
    unset https_proxy
    unset no_proxy
    export PS1=$_PROXY_PS1_ORI
    echo "proxy disabled now"
}

function with_proxy() {
    no_proxy="localhost,127.0.0.1,localaddress,.localdomaiin.com" https_proxy=http://127.0.0.1:1080 http_proxy=http://127.0.0.1:1080 all_proxy=socks5://127.0.0.1:1080  "$@"
}

# Get my ip
myip() {
    curl epurs.com/ip
    # dig +short myip.opendns.com @resolver1.opendns.com
}
