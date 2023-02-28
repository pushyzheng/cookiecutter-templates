import subprocess

def run_shell(cmd):
    print(cmd)
    ret = subprocess.call(cmd, shell=True)
    if ret != 0:
        print("'{}' returns {}".format(cmd, ret))
        exit(-1)


if __name__ == '__main__':
    run_shell('go mod tidy')