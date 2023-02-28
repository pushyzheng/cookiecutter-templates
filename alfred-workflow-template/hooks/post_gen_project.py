import subprocess


def run_shell(cmd):
    print(cmd)
    ret = subprocess.call(cmd, shell=True)
    if ret != 0:
        print("'{}' returns {}".format(cmd, ret))
        exit(-1)


if __name__ == '__main__':
    run_shell('go mod tidy')
    run_shell('sh ./build.sh')
    run_shell('./dist/{{ cookiecutter.project_name }} -cmd hello -query Golang')
    run_shell('git init')
    run_shell('git add .')
    run_shell('git commit -m "Initial commit"')
