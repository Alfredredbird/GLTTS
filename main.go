package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strings"
    "time"

    "github.com/fatih/color"
    "github.com/mattn/go-tty"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    reader := bufio.NewReader(os.Stdin)
    fmt.Println(`╔═╗╦  ╔╦╗╔╦╗
║ ╦║   ║  ║ 
╚═╝╩═╝ ╩  ╩ `)
    fmt.Println("Select difficulty: Beginner, Intermediate, Expert")
    fmt.Print("Enter choice: ")
    choice, _ := reader.ReadString('\n')
    choice = strings.TrimSpace(strings.ToLower(choice))

    var commands []string

    switch choice {
    case "beginner":
        commands = []string{
            "ls", "pwd", "cd /home", "mkdir new_folder",
            "touch file.txt", "echo Hello", "cd ..", "cd ~",
            "rmdir old_folder", "cat file.txt", "nano file.txt",
            "head file.txt", "tail file.txt", "cp file.txt backup.txt",
            "mv file.txt new_file.txt", "rm file.txt", "clear",
            "date", "whoami", "hostname", "uptime", "df -h",
            "du -h", "ping -c 1 google.com", "ps", "kill 1",
            "id", "groups", "env", "history", "alias ll='ls -la'",
            "unalias ll", "echo $HOME", "man ls", "exit", "sleep 1",
            "uptime", "ls -lh",
            "basename /home/user/file.txt", "dirname /home/user/file.txt",
            "sleep 2", "clear; ls", "echo $PATH", "printf 'Hello\nWorld'",
            "yes | head -n 5", "uptime -p", "cal", "who",
            "w", "date +%F", "date +%T", "echo $$", "echo $?",
            "touch a b c", "ls a b c", "rm a b c", "mkdir test1 test2",
            "rmdir test1 test2", "cd /tmp", "pwd", "ls -1",
            "cp file.txt /tmp/", "mv file.txt /tmp/", "rm -i file.txt",
            "head -n 5 file.txt", "tail -n 5 file.txt", "echo 'Test' >> file.txt",
            "cat file.txt | grep Test", "sort file.txt", "uniq file.txt",
            "clear && date", "uptime -s", "hostnamectl", "whoami",
            "id -u", "id -g", "groups $USER", "env | grep PATH",
            "history | tail -n 10",
        }
    case "intermediate":
        commands = []string{
            "ls -la", "rm -rf temp", "cp file1 file2",
            "mv file1 file2", "cat file.txt", "echo Hello World",
            "chmod 644 file.txt", "chown user:user file.txt", "grep Hello file.txt",
            "find . -name '*.txt'", "tar -czvf archive.tar.gz folder/",
            "tar -xzvf archive.tar.gz", "diff file1 file2", "sort file.txt",
            "uniq file.txt", "head -n 10 file.txt", "tail -n 10 file.txt",
            "wc -l file.txt", "cut -d':' -f1 file.txt", "awk '{print $1}' file.txt",
            "sed 's/old/new/g' file.txt", "ps aux | grep nginx", "kill -9 1234",
            "ping -c 4 8.8.8.8", "scp file.txt user@remote:/path",
            "ssh user@remote", "wget https://example.com/file", "curl https://example.com",
            "nano file.txt", "vim file.txt", "history | grep ls", "alias ll='ls -la'",
            "unalias ll", "df -h", "du -sh folder", "top", "htop", "service nginx restart",
            "find /var/log -type f", "grep -r ERROR /var/log", "cut -d' ' -f1 file.txt",
            "awk '{print $2,$3}' file.txt", "sed -n '1,5p' file.txt", "tar -tvf archive.tar.gz",
            "zip -r archive.zip folder", "unzip archive.zip", "rsync -av /src /dst",
            "scp -r folder user@remote:/path", "ssh -p 2222 user@host",
            "curl -I https://example.com", "wget -qO- https://example.com",
            "netstat -tuln", "ss -tuln", "systemctl start nginx", "systemctl stop nginx",
            "systemctl enable nginx", "systemctl disable nginx", "journalctl -u nginx",
            "crontab -l", "crontab -e", "chmod +x script.sh", "./script.sh",
            "tail -f /var/log/syslog", "less file.txt", "more file.txt", "echo $RANDOM",
            "sleep 5 && echo Done", "env | sort", "printenv PATH", "uname -a",
            "uname -r", "df -i", "du -a | sort -n", "watch -n 1 ls",
            "history | tail -n 20", "uptime | awk '{print $3}'", "ping -c 2 yahoo.com",
            "dig example.com", "nslookup example.com",
        }
    case "expert":
        commands = []string{
            "find . -type f -name \"*.txt\"",
            "tar -czvf archive.tar.gz folder/",
            "grep -r \"pattern\" /home/user",
            "chmod -R 755 /var/www/html",
            "sudo systemctl restart nginx",
            "iptables -L", "iptables -A INPUT -p tcp --dport 22 -j ACCEPT",
            "systemctl status apache2", "journalctl -xe", "df -hT",
            "du -sh /var/log/*", "netstat -tulnp", "ss -tulnp", "lsof -i :80",
            "strace ls", "tcpdump -i eth0", "mount | grep /dev/sda1", "umount /mnt/usb",
            "rsync -avz /src /dst", "docker ps -a", "docker run -it ubuntu bash",
            "kubectl get pods", "kubectl describe pod podname", "git clone https://repo.git",
            "git pull origin main", "git commit -am 'message'", "git push origin main",
            "sed -i 's/old/new/g' file.txt", "awk '{sum+=$1} END {print sum}' file.txt",
            "openssl genrsa -out key.pem 2048", "openssl req -new -key key.pem -out csr.pem",
            "chroot /mnt", "dd if=/dev/zero of=file.img bs=1M count=100", "parted /dev/sda",
            "mount -o loop file.img /mnt", "truncate -s 1G file.img", "nc -lvp 1234",
            "nmap -sV 192.168.1.1", "rsyslogd -n", "crontab -e", "systemctl daemon-reload",
            "tcpdump -i eth0 port 22", "iptables -F", "iptables -X", "systemctl list-unit-files",
            "systemctl list-units --failed", "journalctl -f", "journalctl --since '1 hour ago'",
            "strace -p 1234", "lsof -nP | grep LISTEN", "watch -n 2 df -h",
            "nc -zv 192.168.1.1 80", "dig +short example.com", "nslookup example.com",
            "git branch -a", "git checkout -b new-branch", "git merge main", "git rebase main",
            "docker exec -it container bash", "docker logs container", "docker build -t img .",
            "kubectl apply -f deployment.yaml", "kubectl logs podname", "kubectl exec -it podname -- bash",
            "openssl enc -aes-256-cbc -in file.txt -out file.enc", "openssl dgst -sha256 file.txt",
            "dd if=/dev/sda of=disk.img bs=1M", "parted /dev/sda print", "mount -t ext4 /dev/sdb1 /mnt",
            "umount /mnt/usb", "rsync -av --delete /src /dst", "crontab -l | grep backup",
            "systemctl restart sshd", "systemctl stop apache2", "journalctl -u sshd", "tcpdump -i any icmp",
            "iptables -A OUTPUT -p tcp --dport 443 -j ACCEPT", "ss -s", "lsof /var/log/syslog",
            "strace -e trace=open,read,write ls", "nc -lvnp 8080",
        }
    default:
        fmt.Println("Invalid choice, defaulting to Beginner")
        commands = []string{
            "ls", "pwd", "cd /home", "mkdir new_folder",
            "touch file.txt", "echo Hello",
            "cd ..", "cd ~", "rmdir old_folder", "cat file.txt",
            "nano file.txt", "head file.txt", "tail file.txt", "cp file.txt backup.txt",
            "mv file.txt new_file.txt", "rm file.txt", "clear", "date", "whoami",
            "hostname", "uptime", "df -h", "du -h", "ping -c 1 google.com",
            "ps", "kill 1", "id", "groups", "env", "history",
        }
    }

    tty, err := tty.Open()
    if err != nil {
        panic(err)
    }
    defer tty.Close()

    var lastTime float64
    var lastAccuracy float64
    var lastWPM float64

    for {
        target := commands[rand.Intn(len(commands))]

        // Cllear screen
        fmt.Print("\033[H\033[2J")

        // Display last stats if available
        if lastTime > 0 {
            fmt.Printf("Last Command Stats:\n")
            fmt.Printf("Time: %.2fs | Accuracy: %.2f%% | WPM: %.2f\n\n", lastTime, lastAccuracy, lastWPM)
        }

        fmt.Println("Type this command as fast as you can:")

        // color the target command
        color.New(color.FgCyan, color.Bold).Printf("%s\n\n", target)

        fmt.Print("Your input: ")

        input := []rune{}
        start := time.Now()

        for {
            r, err := tty.ReadRune()
            if err != nil {
                panic(err)
            }

            if r == 13 || r == 10 { // enter key
                break
            } else if r == 127 || r == 8 { // Backspace
                if len(input) > 0 {
                    input = input[:len(input)-1]
                }
            } else {
                input = append(input, r)
            }

            // Clear line
            fmt.Print("\r\033[KYour input: ")


            for i := 0; i < len(input); i++ {
                if i < len(target) && input[i] == rune(target[i]) {
                    color.New(color.FgGreen).Printf("%c", input[i])
                } else {
                    color.New(color.FgRed).Printf("%c", input[i])
                }
            }
        }

        elapsed := time.Since(start)

        // Calculate accuracy
        correct := 0
        for i := 0; i < len(input); i++ {
            if i < len(target) && input[i] == rune(target[i]) {
                correct++
            }
        }

        accuracy := float64(correct) / float64(len(target)) * 100
        wpm := float64(len(input)) / 5.0 / elapsed.Minutes()

        // Save last stats
        lastTime = elapsed.Seconds()
        lastAccuracy = accuracy
        lastWPM = wpm

        // Show current round stats
        fmt.Println()
        fmt.Printf("\nTime: %.2fs\n", lastTime)
        fmt.Printf("Accuracy: %.2f%%\n", lastAccuracy)
        fmt.Printf("WPM: %.2f\n", lastWPM)

        fmt.Print("\nPress Enter for next command...")
        tty.ReadRune() 
    }
}
