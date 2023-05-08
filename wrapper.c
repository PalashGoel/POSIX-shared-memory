#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <fcntl.h>
#include <sys/shm.h>
#include <sys/stat.h>
#include <sys/mman.h>
#include <unistd.h>

int my_shm_open(char* filename, int SIZE){
    int shm_fd;
    shm_fd = shm_open(filename,O_CREAT | O_RDWR , 0666);
    if(shm_fd == -1){
        return -1;
    }
    return shm_fd;
}

int my_shm_open_read(char* filename, int SIZE){
    int shm_fd;
    shm_fd = shm_open(filename,O_RDONLY, 0666);
    if(shm_fd == -1){
        return -1;
    }
    return shm_fd;
}


int my_shm_ftruncate(int shm_fd, int SIZE){
    int val;
    val = ftruncate(shm_fd, SIZE);
    return val ;
}

void my_shm_map_write(int shm_fd, int SIZE,char* message){
    void* ptr;
    ptr = mmap(0, SIZE, PROT_WRITE, MAP_SHARED, shm_fd, 0); 

    sprintf(ptr, "%s", message);
    ptr += strlen(message);
    return ;
}


char* my_shm_map_read(int shm_fd, int SIZE){
    void* ptr;
    ptr = mmap(0, SIZE, PROT_READ, MAP_SHARED, shm_fd, 0);

    //printf("%s", (char*)ptr);
    return (char*)ptr;
}

void my_shm_unlink(char* filename){
    shm_unlink(filename);
    return ;
}

